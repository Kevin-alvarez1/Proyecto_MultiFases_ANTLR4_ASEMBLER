mov x0, #0

mov x0, #0

.data
float_val_var1_0: .float 10.000000
msg_0: .ascii "==== b1 ===\n"
msg_1: .ascii "0\n"
msg_2: .ascii "==== b2 ===\n"
msg_3: .ascii "false\n"
msg_4: .ascii "==== global ===\n"
msg_5: .ascii "10.000000\n"

.text
ldr s0, =float_val_var1_0
mov x0, #1
adr x1, msg_0
mov x2, #12
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #12
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_4
mov x2, #16
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_5
mov x2, #10
mov w8, #64
svc #0