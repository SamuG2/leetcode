class Solution:
    def validPath(self, n: int, edges: list[list[int]], source: int, destination: int) -> bool:
        stack = set()
        i=0
        while i<len(edges):
            if source in edges[i]:
                for e in edges[i]:
                    if e == destination:
                        return True
                    if e != source:
                        stack.add(e)
                edges.pop(i)
            else:
                i+=1
        for sor in stack:
            if self.validPath(n, edges, sor, destination):
                return True
        return False        
            
        # edges_dic = {n: edges[n] for n in edges}
        # found = False
        # stack = [source]
        # while stack:
        #     for edge in edges[stack[0]]:
        #         if edge == destination:
        #             return True
        #         if edges_dic.get(stack[0], False):
    
    
                    
