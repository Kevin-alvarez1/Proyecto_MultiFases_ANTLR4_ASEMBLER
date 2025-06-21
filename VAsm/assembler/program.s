.data
float_val_var211_11: .float 1.120000
String_val_var411_11: .ascii "Beta de la beta"
msg_0: .ascii "=== Declaracion sin tipo ===\n"
msg_1: .ascii "100\n"
msg_2: .ascii "1.120000\n"
msg_3: .ascii "false\n"
msg_4: .ascii "Beta de la beta\n"
msg_5: .ascii "=== Asignaciones Declaracion sin valor ===\n"
float_val_var211_12: .float 1.125600
String_val_var411_12: .ascii "Prueba 3"
msg_6: .ascii "1155\n"
msg_7: .ascii "1.125600\n"
msg_8: .ascii "Prueba 3\n"

.global _start
.text
_start:
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
mov x2, #29
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #4
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #9
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
mov x2, #43
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_6
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_7
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_8
mov x2, #9
mov w8, #64
svc #0

    ldp x29, x30, [sp], #16
    ret