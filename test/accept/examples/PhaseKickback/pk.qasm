OPENQASM 2.0;
include "qelib1.inc";

qreg reg1[2];
qreg reg2[1];

h reg1[0];
h reg1[1];
cu1(0.785398163397448) reg1[0],reg2[0];
cu1(1.57079632679490) reg1[1],reg2[0];
