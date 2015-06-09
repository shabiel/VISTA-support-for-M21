# M21 Support
The routines and globals here are a preliminary effort to support M21 on VISTA. Most of VISTA works right now, with a few exceptions.

## Install instructions
You need to pack the routines into a routine archive (see PackRO.py in the OSEHRA VistA repo), then import them into M21 with VISTA installed via %RI.

Import the globals via %GI.

`D ^ZTMGRSET`, and choose `M21`. Make sure you select the correct UCI and VOL for your current M21 instance. Unlike in GT.M it really matters here.

`D ^DINIT` and choose `M21`.

## What works and what doesn't work
The simple stuff all works (anything that does a dumb terminal write; e.g. menu man).

Screenman works 100%.

Listman works with the exception of the arrow and page keys. These need to be supported explicitly in the XGF* routines for M21. I didn't bother.

Error trap works 100%.

Taskman works 100%. There are a couple of bugs in ZTLOAD1: The transactions have been removed (this is an already known bug); and the $GET on $INCREMENT support is flipped: the second arg has been changed from 1 to 0. 

TCP connections have a problem. There is no flush buffer instruction that explicitly flushes the buffer and does nothing else, like in Cache and GT.M. This necessitated a workaround in XWBRB, which I am not including here. I am hoping that this will get fixed at some point.

You have to set TCP connections to TCP_NODELAY. I haven't figured out how to do that.

I managed to get CPRS running with Christopher Edwards' help.

The device manager set-up is tricky. I wasn't able to get ^%ZIS to work properly, but I managed to set-up the important devices.

Home devices need to be set-up with a $I of 1, 2, 3 etc. You must create one for every expected process.

Null devices have to be set-up with a $I of 8050.

```
Select DEVICE NAME: NULL      BIT BUCKET     8050     
ANOTHER ONE: 
STANDARD CAPTIONED OUTPUT? Yes//   (Yes)
Include COMPUTED fields:  (N/Y/R/B): NO//  - No record number (IEN), no Computed Fields

NAME: NULL                              $I: 8050                                ASK DEVICE: NO
  ASK PARAMETERS: NO                    SIGN-ON/SYSTEM DEVICE: YES              LOCATION OF TERMINAL: BIT BUCKET
  OPEN COUNT: 83527                     SUBTYPE: P-DEC                          TYPE: TERMINAL



Select DEVICE NAME: 2  M21 CONSOLE 1    ptys     2     
ANOTHER ONE: 
STANDARD CAPTIONED OUTPUT? Yes//   (Yes)
Include COMPUTED fields:  (N/Y/R/B): NO//  - No record number (IEN), no Computed Fields

NAME: M21 CONSOLE 1                     $I: 2                                   SIGN-ON/SYSTEM DEVICE: YES
  LOCATION OF TERMINAL: ptys            OPEN COUNT: 70                          SUBTYPE: C-VT320/132
  TYPE: TERMINAL
```
