.data
var1: .quad 0
var2: .quad 0
resultado_suma_0: .quad 0
msg_0: .ascii "110\n"
a: .quad 0
b: .quad 0
resultado_suma_1: .quad 0
msg_1: .ascii "5\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #4
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

// Suma 1
adr x10, var1
ldr x10, [x10]
adr x11, var2
ldr x11, [x11]
add x12, x10, x11
adr x13, resultado_suma_0
str x12, [x13]


// Suma 2
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
add x12, x10, x11
adr x13, resultado_suma_1
str x12, [x13]