class Solution:
    def numWays(self, words: List[str], target: str) -> int:
        import collections

        mapp = collections.defaultdict(collections.Counter)
        for word in words:
            for char_ind, char in enumerate(word):
                mapp[char_ind][char] += 1

        @lru_cache(None)
        def dp(it, iw): 
            nonlocal target
            if it == len(target):
                return 1
            if iw == len(words[0]):
                return 0

            if target[it] in mapp[iw]:
                r1 = dp(it + 1, iw + 1)  
                r2 = dp(it, iw + 1)  
                return r1 * mapp[iw][target[it]] + r2
            else:
                r1 = dp(it, iw + 1)
                return r1

        mod = 10 ** 9 + 7
        return dp(0, 0) % mod