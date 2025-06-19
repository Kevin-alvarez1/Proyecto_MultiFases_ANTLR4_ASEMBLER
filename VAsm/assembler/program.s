mov x0, #1

.data
msg_0: .ascii "1\n"

.text
mov x0, #1
adr x1, msg_0
mov x2, #2
mov w8, #64
svc #0
