package hash

import (
    "hash/fnv"
)

func Hash(data []byte, seed uint32) uint32 {
    h := fnv.New32a()
    h.Write(data)
    return h.Sum32() + seed
}
