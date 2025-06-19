mov x0, #112

mov x0, #1

mov x0, #112

mov x0, #1

mov x0, #0

mov x0, #0

.data
var1: .ascii "prueba"
float_val_var3: .float 1.100000
var11: .ascii "prueba"
float_val_var31: .float 1.100000
var12: .ascii ""
float_val_var32: .float 0.000000
msg_0: .ascii "=== declaracion sin tipo ===\n"
msg_1: .ascii "prueba\n"
msg_2: .ascii "112\n"
msg_3: .ascii "1.100000\n"
msg_4: .ascii "true\n"
msg_5: .ascii "=== declaracion con tipo y valor===\n"
msg_6: .ascii "prueba\n"
msg_7: .ascii "112\n"
msg_8: .ascii "1.100000\n"
msg_9: .ascii "true\n"
msg_10: .ascii "=== declaracion con tipo sin valor===\n"
msg_11: .ascii "\n"
msg_12: .ascii "0\n"
msg_13: .ascii "0.000000\n"
msg_14: .ascii "false\n"

.text
ldr s0, =float_val_var3
ldr s0, =float_val_var31
ldr s0, =float_val_var32
mov x0, #1
adr x1, msg_0
mov x2, #29
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #7
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #4
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_4
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_5
mov x2, #36
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_6
mov x2, #7
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_7
mov x2, #4
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_8
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_9
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_10
mov x2, #38
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_11
mov x2, #1
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_12
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_13
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_14
mov x2, #6
mov w8, #64
svc #0
