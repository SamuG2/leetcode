class Solution:
#     def maxPalindromes(self, s: str, k: int) -> int:
#         stack = [[s, k]]
#         res = 0
        
#         while stack:
#             i=0
#             while i<=len(stack[0][0])-stack[0][1]:
#                 if self.isPalindrome(stack[0][0][i:i+stack[0][1]]):
#                     res+=1
#                     print(stack[0][0][i:i+stack[0][1]])
#                     stack.append([stack[0][0][i+stack[0][1]:], stack[0][1]])
#                     stack.append([stack[0][0][:i], stack[0][1]+1])
#                     stack.pop(0)
#                     i=0
#                 else:
#                     i+=1
#             if stack[0][1] > len(stack[0][0]):
#                 stack.pop(0)
#             else:
#                 stack[0][1]+=1
#         return res
                

#     def isPalindrome(self, s: str) -> bool:
#         for i in range(len(s)//2):
#             if s[i] != s[len(s)-i-1]:
#                 return False
#         return True


    def maxPalindromes(self, S: str, k: int) -> int:
        N, ans, start = len(S), 0, 0
        for center in range(2 * N - 1):
            left = center // 2
            right = left + center % 2
            while left >= start and right < N and S[left] == S[right]:
                if right + 1 - left >= k: 
                    ans += 1
                    start = right + 1
                    break
                left -= 1
                right += 1
        return ans


if __name__ == "__main__":
    #s = "abaccdbbd"
    s = "nidinhplkemyryyrymeklphnidinpwlkogggifpupxmxsxxsxmxpupfigggoklwpkxrrcrbytyjqbpbqidvwpymvgygvmypwvdiqbpbqjytybrcrrajrfvgwzmniukskuinmzwgvfremvyhunljhjbuszfozofzsubjhjlnuhyvmesodgkgvnyeobvvbbvvboeynvbkrmkkkkmrkbmhseoqhphkpimegfaelmwlibbbbilwmleafgemipkhphqoeshmehidggatdaeevevwqqlhncxlcpkvqnnqvkpclxcnhlqqwveveeadtaggdojcryfdckcuwwinimpxybthtbyxpminiwwuckcdfyrcjgxkvsibhaodwppwdoahbisvkddquufokwjdzzvckybmwxmbjbmxwmbykcvzzdjwkofuuqsbmmfbxaynonlxkzuhhlqwwqlhhuzkxlnonyaxbfmmbszfrrfzlztimifvywnjwidebfegjixugqgeuuegqguxijgefbediwjnwyvfimitrojxfbglwxxhmlctqkrorkqtclmhxxwlgbfxjoxkaccdajnuegrytenttnetyrgeunjadccahsammbxbbxbmmashfdznohifexneahpimkkmiphaenxejoinhnnzpyuegclzerejejerezlcgeuypznnhniojiwgrixqfvmzdjjdzmvfqxirgwxcnbubnccnxxncskpkbqbkyqjsclpskkreasklbboobblksaerkksplcsjkklqqjwjqqlkkpjuwslpeomxxmoeplswtonqgncvfazkoguviuiiuivugokzafvcngqikluwzutwkrlqmoegypztwhpdwrwdphwtzpygeomqlrkwtuzwulktelwojlfoxphulqugduxkkkkxudguqluhpxophkdxnlifuvgcbbcgvufilnjrokcphagutlciauhujdcqqztushsutzqqcdjuhuaicltugalusowprahtdddlgafxssevondrvxvrdnovessxfagldddtharpwocnpensowmbbdctsnbgcgioarfpwwpfraoigcgbnstcdbbmwosegscreyeedzkuoacfjpsekkespjfcaoukzdeeyercsgupmxvidfqwqmpkfkpmqwqfdivxmlmsekutsejvppvjestukesnfmrgcttcgrmfapnwvjebieeweeibejvwnvk"
    k = 4
    
    print(Solution().maxPalindromes(s, k))