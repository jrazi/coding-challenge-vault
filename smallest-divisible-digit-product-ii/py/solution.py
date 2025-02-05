# https://leetcode.com/problems/smallest-divisible-digit-product-ii/

class Solution:
    def smallestNumber(self, number: str, target: int) -> str:
        DIGIT_FACTORS = [
            (0, 0, 0, 0),
            (0, 0, 0, 0),
            (1, 0, 0, 0),
            (0, 1, 0, 0),
            (2, 0, 0, 0),
            (0, 0, 1, 0),
            (1, 1, 0, 0),
            (0, 0, 0, 1),
            (3, 0, 0, 0),
            (0, 2, 0, 0),
        ]

        def factorize(n):
            counts = [0]*4
            primes = [2,3,5,7]
            for i in range(4):
                while n % primes[i] == 0:
                    counts[i] += 1
                    n //= primes[i]
            return counts, n

        req, rem = factorize(target)
        if rem != 1:
            return "-1"
        
        a_req, b_req, c_req, d_req = req
        INF = 10**18
        dp = [[[[INF]*(d_req+1) for _ in range(c_req+1)] for __ in range(b_req+1)] for ___ in range(a_req+1)]
        dp[0][0][0][0] = 0

        for a in range(a_req+1):
            for b in range(b_req+1):
                for c in range(c_req+1):
                    for d in range(d_req+1):
                        if (a,b,c,d) == (0,0,0,0):
                            continue
                        best = INF
                        for dig in range(1,10):
                            da,db,dc,dd = DIGIT_FACTORS[dig]
                            pa = max(a-da,0)
                            pb = max(b-db,0)
                            pc = max(c-dc,0)
                            pd = max(d-dd,0)
                            best = min(best, dp[pa][pb][pc][pd]+1)
                        dp[a][b][c][d] = best

        def build_suffix(ln, state):
            a,b,c,d = state
            res = []
            for _ in range(ln):
                for dig in range(1,10):
                    da,db,dc,dd = DIGIT_FACTORS[dig]
                    pa = max(a-da,0)
                    pb = max(b-db,0)
                    pc = max(c-dc,0)
                    pd = max(d-dd,0)
                    if dp[pa][pb][pc][pd] <= ln-len(res)-1:
                        res.append(str(dig))
                        a,b,c,d = pa,pb,pc,pd
                        break
            return "".join(res)

        prefix = [(0,0,0,0)]
        valid = [True]
        for ch in number:
            d = int(ch)
            curr = list(prefix[-1])
            if d != 0:
                f = DIGIT_FACTORS[d]
                for i in range(4):
                    curr[i] += f[i]
            prefix.append(tuple(curr))
            valid.append(valid[-1] and (d != 0))

        if valid[-1]:
            total = prefix[-1]
            if total[0]>=a_req and total[1]>=b_req and total[2]>=c_req and total[3]>=d_req:
                return number

        for i in range(len(number)-1,-1,-1):
            if not valid[i]:
                continue
            curr = int(number[i])
            for cand in range(curr+1,10):
                df = DIGIT_FACTORS[cand]
                prev = prefix[i]
                new = [prev[j]+df[j] for j in range(4)]
                ra = max(a_req-new[0],0)
                rb = max(b_req-new[1],0)
                rc = max(c_req-new[2],0)
                rd = max(d_req-new[3],0)
                need = len(number)-i-1
                if dp[ra][rb][rc][rd] <= need:
                    suf = build_suffix(need, (ra,rb,rc,rd))
                    return number[:i] + str(cand) + suf

        length = len(number)+1
        while True:
            if dp[a_req][b_req][c_req][d_req] <= length:
                return build_suffix(length, (a_req,b_req,c_req,d_req))
            length += 1
