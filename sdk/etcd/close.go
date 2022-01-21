package etcd

func (e *etcd) Close() error {
	return e.client.Close()
}
