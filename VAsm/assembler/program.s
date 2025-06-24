.data
x: .float 0.0
float_val_x_2: .float 8.900000
y: .float 0.0
float_val_y_3: .float 3.200000
z: .quad 0
msg_0: .ascii "2\n"
msg_1: .ascii "0\n"
a: .quad 0
b: .quad 0
msg_2: .ascii "1\n"

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

mov x0, #1
adr x1, msg_0
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #2
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0