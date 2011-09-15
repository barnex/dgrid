include $(GOROOT)/src/Make.inc

main:*.go
	$(GC) -I_obj *.go
	$(LD) -L_obj -o dgrid *.$O

.PHONY:	clean
clean:
	rm -f  main *.a *.$O $(CLEANFILES)
