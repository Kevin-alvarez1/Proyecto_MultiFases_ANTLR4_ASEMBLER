mov x0, #1

mov x0, #1

.data
float_val_var2: .float 1.100000
var4: .ascii "Ola mundo"
msg_0: .ascii "=== Declaracion explicita ===\n"
msg_1: .ascii "1\n"
msg_2: .ascii "1.100000\n"
msg_3: .ascii "true\n"
msg_4: .ascii "Ola mundo\n"

.text
ldr s0, =float_val_var2
mov x0, #1
adr x1, msg_0
mov x2, #30
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_4
mov x2, #10
mov w8, #64
svc #0