.data
uno: .quad 0
msg_0: .ascii "true\n"
msg_1: .ascii "false\n"
f: .quad 0

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #6
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0

// Operación lógica NOT !uno
adr x10, uno
ldr x10, [x10]
cmp x10, #0
cset x0, eq