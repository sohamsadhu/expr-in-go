cmd = expr
include $(GOROOT)/src/Make.$(GOARCH)
all: $O/$(cmd)
clean:		 ; rm -rf [$(OS)] *.[$(OS)] # these lines do not have any source just the target and command. 
doc:	 	 ; godoc -cmdroot=.. $(cmd)
$O:	   	  ; mkdir $@ #the @ sign is the full name of the current target so making a directory for the current target that could be 6 not making any sense. 
$O/%: $O %.go;	 $(GC) $(GCFLAGS) -o $*.$O $*.go && \
              $(LD) $(LDFLAGS) -o $@ $*.$O
test: all	; $O/$(cmd) 2 \* \( 3 + 4 \) / \( 2 - 4 \)

