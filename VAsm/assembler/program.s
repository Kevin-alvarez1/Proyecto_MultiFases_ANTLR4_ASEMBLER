.data
a: .float 0.0
float_val_a_0: .float 3.140000
b: .float 0.0
float_val_b_1: .float 3.300000
msg_0: .ascii "0\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #2
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0