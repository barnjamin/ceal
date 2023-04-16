#include "avm.hpp"
#include "abi.hpp"

uint64 avm_main()
{

    uint64 val = 1;
    bytes v1 = abi_encode(&val);
    uint64 v2 = *(uint64*)abi_decode(v1);

    avm_assert(val == v2);

    return 1;
}