

# idea 1:creamos un diccionario con 0..n-1 claves. Cada valor otro diccionario. Vamos rellenando cada
# diccionario con las distancias que de a a b, de forma que solo tenemos que iterar una vez
class Solution:
    def sumOfDistancesInTree(self, n: int, edges: list[list[int]]) -> list[int]:
        distances_dic = {i: {i:0} for i in range(n)}
        res = [0 for _ in range(n)]

        for edge in edges:
            distances_dic[edge[0]][edge[1]]=1
            distances_dic[edge[1]][edge[0]]=1
        for i in range(n):
            
            
            
            
