; hi erin here. i think this is stupid their hints are really cryptic augh i dont know how to do anything this is rather daft. i had to ask chatgpt to explain things because they did a very poor job of it

; Everything that comes after a semicolon (;) is a comment

; Assembler-time constants may be defined using 'equ'

section .text

; You should implement functions in the .text section

; the global directive makes a function visible to the test files
global expected_minutes_in_oven
expected_minutes_in_oven:
    mov rax, 40
    ret

global remaining_minutes_in_oven
remaining_minutes_in_oven:
    call expected_minutes_in_oven
    sub rax, 25
    ret

global preparation_time_in_minutes
preparation_time_in_minutes:
    mov rax, rdi
    imul rax, 2
    ret

global elapsed_time_in_minutes
elapsed_time_in_minutes:
    mov rax, rdi
    call preparation_time_in_minutes
    add rax, rsi
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
