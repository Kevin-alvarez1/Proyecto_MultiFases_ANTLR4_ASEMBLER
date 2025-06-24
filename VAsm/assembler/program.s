.data
a: .quad 0
b: .quad 0
f1: .float 0.0
float_val_f1_8: .float 2.500000
f2: .float 0.0
float_val_f2_9: .float 5.000000
resultado_suma_0: .quad 0
msg_0: .ascii "13\n"
resultado_suma_1: .float 0.0
msg_1: .ascii "12.500000\n"
resultado_suma_2: .float 0.0
msg_2: .ascii "7.500000\n"
resultado_resta_3: .float 0.0
msg_3: .ascii "7.000000\n"
resultado_resta_4: .float 0.0
resultado_resta_5: .float 0.0
msg_4: .ascii "2.000000\n"
resultado_mul_6: .float 0.0
msg_5: .ascii "30.000000\n"
resultado_mul_7: .float 0.0
msg_6: .ascii "25.000000\n"
resultado_mul_8: .float 0.0
resultado_div_9: .float 0.0
msg_7: .ascii "3.333333\n"
resultado_div_10: .float 0.0
msg_8: .ascii "4.000000\n"
resultado_div_11: .float 0.0
msg_9: .ascii "0.500000\n"
msg_10: .ascii "<nil>\n"
resultado_suma_12: .float 0.0
resultado_resta_13: .float 0.0
msg_11: .ascii "-7.500000\n"
resultado_mul_14: .float 0.0
resultado_div_15: .float 0.0
msg_12: .ascii "1.666667\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #3
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_4
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_5
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_6
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_7
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_8
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_9
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_10
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_11
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_6
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_12
mov x2, #9
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0


// Suma 1
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
add x12, x10, x11
adr x13, resultado_suma_0
str x12, [x13]


// Suma 2
adr x10, a
ldr s0, [x10]
adr x11, f1
ldr s1, [x11]
fadd s2, s0, s1
adr x12, resultado_suma_1
str s2, [x12]


// Suma 3
adr x10, f1
ldr s0, [x10]
adr x11, f2
ldr s1, [x11]
fadd s2, s0, s1
adr x12, resultado_suma_2
str s2, [x12]


// Resta 4
adr x10, a
ldr s0, [x10]
adr x11, b
ldr s1, [x11]
fsub s2, s0, s1
adr x12, resultado_resta_3
str s2, [x12]


// Resta 5
adr x10, a
ldr s0, [x10]
adr x11, f1
ldr s1, [x11]
fsub s2, s0, s1
adr x12, resultado_resta_4
str s2, [x12]


// Resta 6
adr x10, f2
ldr s0, [x10]
adr x11, b
ldr s1, [x11]
fsub s2, s0, s1
adr x12, resultado_resta_5
str s2, [x12]


// Multiplicación 7
adr x10, a
ldr s0, [x10]
adr x11, b
ldr s1, [x11]
fmul s2, s0, s1
adr x12, resultado_mul_6
str s2, [x12]


// Multiplicación 8
adr x10, a
ldr s0, [x10]
adr x11, f1
ldr s1, [x11]
fmul s2, s0, s1
adr x12, resultado_mul_7
str s2, [x12]


// Multiplicación 9
adr x10, f1
ldr s0, [x10]
adr x11, f2
ldr s1, [x11]
fmul s2, s0, s1
adr x12, resultado_mul_8
str s2, [x12]


// División 10
adr x10, a
ldr s0, [x10]
adr x11, b
ldr s1, [x11]
fdiv s2, s0, s1
adr x12, resultado_div_9
str s2, [x12]


// División 11
adr x10, a
ldr s0, [x10]
adr x11, f1
ldr s1, [x11]
fdiv s2, s0, s1
adr x12, resultado_div_10
str s2, [x12]


// División 12
adr x10, f2
ldr s0, [x10]
adr x11, a
ldr s1, [x11]
fdiv s2, s0, s1
adr x12, resultado_div_11
str s2, [x12]


// Suma 13
adr x10, f1
ldr s0, [x10]
adr x11, a
ldr s1, [x11]
fadd s2, s0, s1
adr x12, resultado_suma_12
str s2, [x12]


// Resta 14
adr x10, f1
ldr s0, [x10]
adr x11, a
ldr s1, [x11]
fsub s2, s0, s1
adr x12, resultado_resta_13
str s2, [x12]


// Multiplicación 15
adr x10, f1
ldr s0, [x10]
adr x11, a
ldr s1, [x11]
fmul s2, s0, s1
adr x12, resultado_mul_14
str s2, [x12]


// División 16
adr x10, f2
ldr s0, [x10]
adr x11, b
ldr s1, [x11]
fdiv s2, s0, s1
adr x12, resultado_div_15
str s2, [x12]