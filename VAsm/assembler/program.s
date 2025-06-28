.data
puntos_while: .quad 0
i: .quad 0
suma1: .quad 0
uno: .quad 0
cinco: .quad 0
diez: .quad 0
msg_0: .ascii "0\n"
msg_1: .ascii "1\n"
msg_2: .ascii "2\n"
msg_3: .ascii "3\n"
msg_4: .ascii "4\n"
msg_5: .ascii "OK suma1 == 10\n"
msg_6: .ascii "OK i == 5\n"

.global _start
.text
_start:

// ====== FOR CONDICIONAL ======
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
adr x1, msg_2
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_4
mov x2, #2
mov w8, #64
svc #0


//  ====== IF ====== 1
mov x0, #1
adr x1, msg_5
mov x2, #15
mov w8, #64
svc #0

if_1_else:

if_1_end:

//  ====== IF ====== 2
mov x0, #1
adr x1, msg_6
mov x2, #10
mov w8, #64
svc #0

if_2_else:

if_2_end:
mov x0, #1
adr x1, msg_2
mov x2, #2
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0

// Comparación < entre i y cinco
adr x10, i
ldr x10, [x10]
adr x11, cinco
ldr x11, [x11]
cmp x10, x11
cset x0, lt
// Comparación < entre i y cinco
adr x10, i
ldr x10, [x10]
adr x11, cinco
ldr x11, [x11]
cmp x10, x11
cset x0, lt
// Comparación < entre i y cinco
adr x10, i
ldr x10, [x10]
adr x11, cinco
ldr x11, [x11]
cmp x10, x11
cset x0, lt
// Comparación < entre i y cinco
adr x10, i
ldr x10, [x10]
adr x11, cinco
ldr x11, [x11]
cmp x10, x11
cset x0, lt
// Comparación < entre i y cinco
adr x10, i
ldr x10, [x10]
adr x11, cinco
ldr x11, [x11]
cmp x10, x11
cset x0, lt
// Comparación < entre i y cinco
adr x10, i
ldr x10, [x10]
adr x11, cinco
ldr x11, [x11]
cmp x10, x11
cset x0, lt
// Comparación == entre suma1 y diez
adr x10, suma1
ldr x10, [x10]
adr x11, diez
ldr x11, [x11]
cmp x10, x11
cset x0, eq
// Comparación == entre i y cinco
adr x10, i
ldr x10, [x10]
adr x11, cinco
ldr x11, [x11]
cmp x10, x11
cset x0, eq