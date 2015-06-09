Unencoded M21 Global Save from MGR,SYS on 09-JUN-2015 at 07:01 AM
M21 FM OS node
^DD("OS",101,0)
M21^^32000^32000^^1^255
^DD("OS",101,1)
B
^DD("OS",101,8)
X ^DD("$O")
^DD("OS",101,18)
I $D(^ (X))
^DD("OS",101,"DEL")
X "ZR  ZS @X" K ^UTILITY("%RD",X)
^DD("OS",101,"EOFF")
U $I:(::::1)
^DD("OS",101,"EON")
U $I:(:::::1)
^DD("OS",101,"LOAD")
S %N=0 X "ZL @X F XCNP=XCNP+1:1 S %N=%N+1,%=$T(+%N) Q:$L(%)=0  S @(DIF_XCNP_"",0)"")=%"
^DD("OS",101,"NO-TYPE-AHEAD")
U $I:(::::100663296)
^DD("OS",101,"RM")
U:IOT["TRM" $I:X
^DD("OS",101,"RSEL")
K ^UTILITY($J) G ^%RSEL
^DD("OS",101,"SDP")
O @("DIO:"_DLP) F %=0:0 U DIO R % Q:$ZA=X&($ZB>Y)!($ZA>X)  U IO W:$A(%)'=12 ! W %
^DD("OS",101,"SDPEND")
S X=$ZA,Y=$ZB
^DD("OS",101,"TRMOFF")
U $I:(::::::::$C(13,27))
^DD("OS",101,"TRMON")
U $I:(::::::::$C(0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,127))
^DD("OS",101,"TRMRD")
S Y=$ZB
^DD("OS",101,"TYPE-AHEAD")
U $I:(::::67108864:33554432)
^DD("OS",101,"UCICHECK")
S Y=$$UCICHECK^DINVMSM(X)
^DD("OS",101,"XY")
S $X=IOX,$Y=IOY
^DD("OS",101,"ZS")
ZR  X "S %Y=0 F  S %Y=$O(^UTILITY($J,0,%Y)) Q:%Y=""""  Q:'$D(^(%Y))  ZI ^(%Y)" ZS @X


