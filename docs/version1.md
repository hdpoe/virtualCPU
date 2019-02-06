# Version 1.0

The first version will implement an abstract concept of a CPU. That operates an idealized Fetch-Decode-Execute cycle.

## Components
The following components are part of the virtual CPU.
* Memory - A contigous array wherein each cell holds a uint16, for 16 bits the first 8 will be op code and the next 8 will be data. This means we will be limited to 256 bytes of memory as that will be all that can be addressed.
* Processor - The processor will be responsbile for executing the fetch decode execute cycle and carrying out the instructions specified.
* Registers: The following abstract registers will be used:
  - Accumalator: The result of each operation will be stored in the Accumalator
  - Memory Address Register: This registers holds the address of the memory that will be written to or fetched from depending on the instruction.
  - Memory Data Register: This register holds the data that is to be written to or has just been fetched from memory
  - Program Counter: This register will be auto incremented after each step. It holds the address of the next instruction to execute.
  - Current Instruction Register: This register holds the instruction that is currently being executed.
* Decoder: This component will determine what to do when an instruction is read.
* State/Flags Register: This will be a register that holds flags that reflect the current state of the register.

## Language
The assembly language used will be an abstract language of incredibly limited scope that supports only the following operations

* load {src}  - Loads data from the memory address indicated into the accumalator.
* save {src} - Saves the value in the accumulator to the specified memory address.
* add {value} - Adds the value or memory address indicated to the value in the accumulator and stores the results in the accumalator.
* sub {src} - Same as above but subtracts.
* mul {src} - Same as above but multiplies.
* div {src} - Divides the value in the accumalator by the value indicated in the memory address or a static value.
* cmp {src} - Compares the value at the given memory address to the value in the accumlator by performing accumlator - src operation and updates the flags registers with the result.
* jmp {flag} {src} - If the given flag in the register is set jumps to the memory address indicated.
