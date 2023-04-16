#include "avm.hpp"

uint64 avm_main(){

    uint64 val = 10;

    asm { 
        int 1
        pop
    };

    return 1;
}