.data
msg_0: .ascii "1\n"
msg_1: .ascii "2\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #2
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0


.global fn_suma
fn_suma:
    stp x29, x30, [sp, #-16]!
    mov x29, sp
    mov sp, x29
    ldp x29, x30, [sp], #16
    ret