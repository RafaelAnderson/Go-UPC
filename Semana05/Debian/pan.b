	switch (t->back) {
	default: Uerror("bad return move");
	case  0: goto R999; /* nothing to undo */

		 /* PROC p */

	case 3: // STATE 2
		;
		now.S = trpt->bup.oval;
		;
		goto R999;

	case 4: // STATE 4
		;
		now.c = trpt->bup.oval;
		;
		goto R999;
;
		;
		
	case 6: // STATE 6
		;
		now.c = trpt->bup.oval;
		;
		goto R999;

	case 7: // STATE 7
		;
		now.S = trpt->bup.oval;
		;
		goto R999;

	case 8: // STATE 11
		;
		p_restor(II);
		;
		;
		goto R999;
	}

