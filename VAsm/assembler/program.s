.data
msg_0: .ascii "dentro de main: \n"
msg_1: .ascii "20\n"
msg_2: .ascii "fuera de todo: \n"
msg_3: .ascii "30\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_2
mov x2, #16
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #3
mov w8, #64
svc #0

    bl fn_main
    mov x0, #0
    mov x8, #93
    svc #0


.global fn_main
fn_main:
    stp x29, x30, [sp, #-16]!
    mov x29, sp
    mov sp, x29
mov x0, #1
adr x1, msg_0
mov x2, #17
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #3
mov w8, #64
svc #0

    ldp x29, x30, [sp], #16
    ret