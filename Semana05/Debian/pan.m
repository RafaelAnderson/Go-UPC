#define rand	pan_rand
#define pthread_equal(a,b)	((a)==(b))
#if defined(HAS_CODE) && defined(VERBOSE)
	#ifdef BFS_PAR
		bfs_printf("Pr: %d Tr: %d\n", II, t->forw);
	#else
		cpu_printf("Pr: %d Tr: %d\n", II, t->forw);
	#endif
#endif
	switch (t->forw) {
	default: Uerror("bad forward move");
	case 0:	/* if without executable clauses */
		continue;
	case 1: /* generic 'goto' or 'skip' */
		IfNotBlocked
		_m = 3; goto P999;
	case 2: /* generic 'else' */
		IfNotBlocked
		if (trpt->o_pm&1) continue;
		_m = 3; goto P999;

		 /* PROC p */
	case 3: // STATE 1 - ejercicio05.pml:11 - [((S>0))] (4:0:1 - 1)
		IfNotBlocked
		reached[0][1] = 1;
		if (!((((int)now.S)>0)))
			continue;
		/* merge: S = (S-1)(0, 2, 4) */
		reached[0][2] = 1;
		(trpt+1)->bup.oval = ((int)now.S);
		now.S = (((int)now.S)-1);
#ifdef VAR_RANGES
		logval("S", ((int)now.S));
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 4: // STATE 4 - ejercicio05.pml:12 - [c = (c+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][4] = 1;
		(trpt+1)->bup.oval = ((int)now.c);
		now.c = (((int)now.c)+1);
#ifdef VAR_RANGES
		logval("c", ((int)now.c));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 5: // STATE 5 - ejercicio05.pml:13 - [assert((c<2))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][5] = 1;
		spin_assert((((int)now.c)<2), "(c<2)", II, tt, t);
		_m = 3; goto P999; /* 0 */
	case 6: // STATE 6 - ejercicio05.pml:14 - [c = (c-1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][6] = 1;
		(trpt+1)->bup.oval = ((int)now.c);
		now.c = (((int)now.c)-1);
#ifdef VAR_RANGES
		logval("c", ((int)now.c));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 7: // STATE 7 - ejercicio05.pml:15 - [S = (S+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][7] = 1;
		(trpt+1)->bup.oval = ((int)now.S);
		now.S = (((int)now.S)+1);
#ifdef VAR_RANGES
		logval("S", ((int)now.S));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 8: // STATE 11 - ejercicio05.pml:17 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[0][11] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */
	case  _T5:	/* np_ */
		if (!((!(trpt->o_pm&4) && !(trpt->tau&128))))
			continue;
		/* else fall through */
	case  _T2:	/* true */
		_m = 3; goto P999;
#undef rand
	}

