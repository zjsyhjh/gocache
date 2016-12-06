## 一致性Hash算法

![consistent-hash](https://github.com/zjsyhjh/gocache/blob/master/png/hash.png?raw=true)

- 如上图所致，每台server复制replicas份映射到圆周上（具体实现中用有序数组实现）
- 一致性hash算法将key分布在如上图所示的圆周上，然后将key存储到离它最近的server中，例如key-2存储在server-2中，key-4存储在server-1中，这样做的好处是可以避免server宕机时带来的映射关系的重大变化（如图所示，如果server-2宕机，则只需重新映射key-2以及key-1就可以，不需要动key-3和key-4，但如果是一般的取模hash映射，则N的大小变化，导致所有的key都需要重新映射）
- 具体实现上，用一个有序数组来存储所有的server以及一个hashmap来存储每个server的具体信息，这里的键值对可能为\<server-hashvalue, server-ip\>
- Add函数用于添加所有的server到数组中，采用32位循环冗余检验和算法得到server的hash值
- Get函数用于获取每个key对应的server信息（例如ip），首先计算key对应的hash值，然后通过二分查找找到最近的server，通过hashmap获取server相应的信息（例如返回server对应的ip）
- go test -v测试全部用例，go test -v -test.run TestHash测试单个函数