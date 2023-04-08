# ceal 👷 EARLY ACCESS 👷

C to Algorand TEAL compiler (transpiler)

AKA use your favorite C++ IDE to write Algorand contracts

## Usage

```cpp
#include <avm.hpp>

void init()
{
	// TODO: initialize app
}

void noop()
{
	// TODO: insert NoOp logic here
}

uint64 avm_main()
{
	if (avm_txn.ApplicationID == 0)
	{
		init();
		return 1;
	}

	if (avm_txn.OnCompletion != NoOp)
	{
		avm_err();
	}

	noop();
}
```
```bash
go run cmd/ceal/main -path contract.cpp
```

## Binaries

See 'dev' tag for the latest build artifacts ->
