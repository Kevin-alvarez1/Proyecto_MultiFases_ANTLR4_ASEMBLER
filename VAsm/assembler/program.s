.data
msg_0: .ascii "223123\n"
msg_1: .ascii "0\n"

.text
mov x0, #1
adr x1, msg_0
mov x2, #7
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #2
mov w8, #64
svc #0
