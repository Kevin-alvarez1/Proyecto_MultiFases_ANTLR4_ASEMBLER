.data
f: .quad 0
c: .quad 0
v: .quad 0
msg_0: .ascii "false\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #6
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0