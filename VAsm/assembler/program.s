.data
a: .quad 0
b: .quad 0
c: .quad 0
text1: .quad 0
String_val_text1_0: .ascii "=== COMPARACION RELACIONALES ==="
text2: .quad 0
String_val_text2_1: .ascii "=== BOOLEANOS DIRECTOS ==="
t: .quad 0
f: .quad 0
msg_0: .ascii "=== COMPARACION RELACIONALES ===\n"
msg_1: .ascii "false\n"
msg_2: .ascii "true\n"
msg_3: .ascii "=== BOOLEANOS DIRECTOS ===\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #33
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #27
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #5
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #6
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #5
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0

// Comparación == entre a y b
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
cmp x10, x11
cset x0, eq
// Comparación == entre a y c
adr x10, a
ldr x10, [x10]
adr x11, c
ldr x11, [x11]
cmp x10, x11
cset x0, eq
// Comparación != entre a y b
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
cmp x10, x11
cset x0, ne
// Comparación > entre b y a
adr x10, b
ldr x10, [x10]
adr x11, a
ldr x11, [x11]
cmp x10, x11
cset x0, gt
// Comparación < entre b y a
adr x10, b
ldr x10, [x10]
adr x11, a
ldr x11, [x11]
cmp x10, x11
cset x0, lt
// Comparación >= entre a y c
adr x10, a
ldr x10, [x10]
adr x11, c
ldr x11, [x11]
cmp x10, x11
cset x0, ge
// Comparación <= entre b y c
adr x10, b
ldr x10, [x10]
adr x11, c
ldr x11, [x11]
cmp x10, x11
cset x0, le
// Operación lógica AND entre t && f
adr x10, t
ldr x10, [x10]
adr x11, f
ldr x11, [x11]
and x12, x10, x11
cmp x12, #0
cset x0, ne
// Operación lógica OR entre t || f
adr x10, t
ldr x10, [x10]
adr x11, f
ldr x11, [x11]
orr x12, x10, x11
cmp x12, #0
cset x0, ne
// Operación lógica OR entre f || f
adr x10, f
ldr x10, [x10]
adr x11, f
ldr x11, [x11]
orr x12, x10, x11
cmp x12, #0
cset x0, ne
// Operación lógica AND entre t && t
adr x10, t
ldr x10, [x10]
adr x11, t
ldr x11, [x11]
and x12, x10, x11
cmp x12, #0
cset x0, ne