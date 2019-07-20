package zk

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"encoding/json"
)

type SdClient struct {
	zkServers []string // 多个节点地址
	zkRoot    string   // 服务根结点
	conn      *zk.Conn // zk的客户端连接
}

func NewClient(zkServers []string, zkRoot string, timeout int) (*SdClient, error) {
	client := new(SdClient)
	client.zkServers = zkServers
	client.zkRoot = zkRoot

	//连接服务器
	conn, _, err := zk.Connect(zkServers, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	client.conn = conn

	//创建服务根结点
	if err := client.ensureRoot(); err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}

func (client *SdClient) ensureRoot() error {
	exists, _, err := client.conn.Exists(client.zkRoot)
	if err != nil {
		return err
	}

	if !exists {
		// Create调用可能会返回节点已存在错误，这是正常现象，因为会存在多进程同时创建节点的可能。如果创建根节点出错，还需要及时关闭连接。
		// 我们不关心节点的权限控制，所以使用zk.WorldACL(zk.PermAll)表示该节点没有权限限制。Create参数中的flag=0表示这是一个持久化的普通节点
		_, err := client.conn.Create(client.zkRoot, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil && err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

func (client *SdClient) Close() {
	client.conn.Close()
}

func (client *SdClient) Register(node *ServiceNode) error {
	if err := client.ensureName(node.Name); err != nil {
		return err
	}

	path := client.zkRoot + "/" + node.Name
	data, err := json.Marshal(node)
	if err != nil {
		return err
	}

	//创建一个保护顺序临时(ProtectedEphemeralSequential)子节点，同时将地址信息存储在节点中。什么叫保护顺序临时节点，首先它是一个临时节点，
	// 会话关闭后节点自动消失。其次，它是个顺序节点，zookeeper自动在名称后面增加自增后缀，确保节点名称的唯一性。同时还是个保护性节点，节点前缀增加了GUID字段，
	// 确保断开重连后临时节点可以和客户端状态对接上.data只是与节点"绑定"的信息，并非节点中真是存在的数据
	_, err = client.conn.CreateProtectedEphemeralSequential(path, data, zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}

	// 节点中存储数据通过此方法设置
	_, err = client.conn.Set(path, data, -1)
	if err != nil {
		return err
	}
	return nil

}
func (client *SdClient) ensureName(name string) error {
	path := client.zkRoot + "/" + name
	exists, _, err := client.conn.Exists(path)
	if err != nil {
		return err
	}

	if !exists {
		_, err := client.conn.Create(path, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil && err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

func (client *SdClient) GetServiceNodeList() (list []string, err error) {
	list, _, err = client.conn.Children(client.zkRoot)
	return list, err
}

func (client *SdClient) GetServiceNodeData(name string) (*ServiceNode, error) {
	if err := client.ensureName(name); err != nil {
		return nil, err
	}

	data, _, err := client.conn.Get(client.zkRoot + "/" + name)
	if err != nil {
		return nil, err
	}

	serviceNode := &ServiceNode{}
	if err = json.Unmarshal(data, serviceNode); err != nil {
		return nil, err
	}

	return serviceNode, nil
}


