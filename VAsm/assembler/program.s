mov x0, #1

mov x0, #1

mov x0, #0

mov x0, #0

mov x0, #100

mov x0, #0

mov x0, #11

mov x0, #0

mov x0, #1112

mov x0, #1

mov x0, #1155

mov x0, #0

.data
float_val_var2_12: .float 1.100000
String_val_var4_12: .ascii "Ola mundo"
msg_0: .ascii "=== Declaracion explicita ===\n"
msg_1: .ascii "1\n"
msg_2: .ascii "1.100000\n"
msg_3: .ascii "true\n"
msg_4: .ascii "Ola mundo\n"
float_val_var21_13: .float 0.000000
String_val_var41_13: .ascii ""
msg_5: .ascii "=== Declaracion sin valor ===\n"
msg_6: .ascii "0\n"
msg_7: .ascii "0.000000\n"
msg_8: .ascii "false\n"
msg_9: .ascii "\n"
float_val_var211_14: .float 1.120000
String_val_var411_14: .ascii "Beta de la beta"
msg_10: .ascii "=== Declaracion sin tipo ===\n"
msg_11: .ascii "100\n"
msg_12: .ascii "1.120000\n"
msg_13: .ascii "false\n"
msg_14: .ascii "Beta de la beta\n"
msg_15: .ascii "=== Asignaciones Declaracion explicita ===\n"
float_val_var2_15: .float 1.120000
String_val_var4_15: .ascii "Prueba 1"
msg_16: .ascii "11\n"
msg_17: .ascii "1.120000\n"
msg_18: .ascii "false\n"
msg_19: .ascii "Prueba 1\n"
msg_20: .ascii "=== Asignaciones Declaracion sin valor ===\n"
float_val_var21_16: .float 1.122100
String_val_var41_16: .ascii "Prueba 2"
msg_21: .ascii "1112\n"
msg_22: .ascii "1.122100\n"
msg_23: .ascii "true\n"
msg_24: .ascii "Prueba 2\n"
msg_25: .ascii "=== Asignaciones Declaracion sin valor ===\n"
float_val_var211_17: .float 1.125600
String_val_var411_17: .ascii "Prueba 3"
msg_26: .ascii "1155\n"
msg_27: .ascii "1.125600\n"
msg_28: .ascii "false\n"
msg_29: .ascii "Prueba 3\n"

.text
ldr s0, =float_val_var2_12
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

ldr s0, =float_val_var21_13
mov x0, #1
adr x1, msg_5
mov x2, #30
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_6
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_7
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_8
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_9
mov x2, #1
mov w8, #64
svc #0

ldr s0, =float_val_var211_14
mov x0, #1
adr x1, msg_10
mov x2, #29
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_11
mov x2, #4
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_12
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_13
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_14
mov x2, #16
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_15
mov x2, #43
mov w8, #64
svc #0

ldr s0, =float_val_var2_15
mov x0, #1
adr x1, msg_16
mov x2, #3
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_17
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_18
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_19
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_20
mov x2, #43
mov w8, #64
svc #0

ldr s0, =float_val_var21_16
mov x0, #1
adr x1, msg_21
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_22
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_23
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_24
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_25
mov x2, #43
mov w8, #64
svc #0

ldr s0, =float_val_var211_17
mov x0, #1
adr x1, msg_26
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_27
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_28
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_29
mov x2, #9
mov w8, #64
svc #0