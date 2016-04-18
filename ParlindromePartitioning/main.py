class Solution(object):
    def partition(self, s):
        """
        :type s: str
        :rtype: List[List[str]]
        """
        l = list(s)
        obj = Solution()
        d = obj.findSameLetter(s)
        l2 = obj.findHuiXingStr(d,l)
        l3 = obj.sortList(l2)
        result = list()
        for i in range(len(l3)):
            l4 = obj.getResult(s,l3[i])
            result.append(l4)
        l5 = list()
        for a in range(len(s)):
            l5.append(s[a])
        result.append(l5)
        return result

        #1:找相同的字母,返回map
    def findSameLetter(self,s):
        l = list(s)
        d = dict()
        for i in range(len(l)):
            if (i + 1) <= len(l):
                key = l[i]
                val = list()
                val.append(i)
                for j in range(i+1 ,len(l)):
                    if(l[j] == l[i]):
                        val.append(j)
                        l[j] = "def"
                d[key] = val
            else:
                pass
        l2 = list()
        for (k,v) in d.items():
            if(len(v) == 1):
                l2.append(k)
        for i in range(len(l2)):
            d.pop(l2[i])
        return d

    #2:从相同的字母中找到回型字符串,返回dict
    def findHuiXingStr(self,d,l):
        l3 = list()
        for (k,v) in d.items():
            l2 = d[k]
            for i in range(len(l2)):
                if i + 1 < len(l2):
                    for j in range(i + 1 , len(l2)):
                        l4 = list()
                        l4.append(l2[i])
                        l4.append(l2[j])
                        obj = Solution()
                        flag = obj.checkHuanxingStr(l,l4)
                        if flag :
                            l3.append(l4)
                        else:
                            continue
        return l3

    #checkHuanxingStr
    def checkHuanxingStr(self,l,l1):
        flag = True
        if(len(l1) != 2):
            print('这个列表的长度应该为2')
        else:
            start = l1[0]
            stop = l1[1]
            # print('checkHuanxingStr{}'.format(l1))
            for i in range(stop - start):
                #print("start:{}--stop{}".format(l[start + i],l[stop - i]))
                if(l[start + i] != l[stop - i]):
                    flag = False
                    break
        return flag

    #3:整合找到的回型字符串
    def sortList(self,l1):
        l2 = list()
        l6 = list()
        for i in range(len(l1)):
            l3 = list()
            l3.append(l1[i])
            l2.append(l3)
            l6.append(l3)

        lNew = list()
        for j in range(len(l1)):
            for k in range(j + 1,len(l1)):
                if l1[k][0] > l1[j][1] or l1[k][1] < l1[j][0]:
                    l4 = list()
                    l4.append(l1[j])
                    l4.append(l1[k])
                    lNew.append(l4)
        l2.extend(lNew)

        l8 = list()
        l8.extend(lNew)
        while len(l8) != 0 :
            l9 = list()
            l9.extend(l8)
            l8 = []
            for a in range(len(l6)):
                for b in range(len(l9)):
                    flag = True
                    for c in range(len(l9[b])):
                        obj = Solution()
                        flag0 = obj.checkL1L2(l9[b][c],l6[a][0])
                        if not flag0:
                            flag = False
                    if flag:
                        l5 = list()
                        for d in range(len(l9[b])):
                            l5.append(l9[b][d])
                        l5.append(l6[a][0])
                        l5 = sorted(l5)
                        if l5 not in l2 :
                            l8.append(l5)
                            l2.append(l5)
        return l2

    def checkL1L2(self,l1,l2):
        if l1[0] > l2[1] or l1[1] < l2[0]:
            return True
        else:
            return False

    def orderList(self,l):
        l = sorted(l)
        l2 = list()
        for i in range(len(l)):
            l2.append(l[i][0])
            l2.append(l[i][1])
        return l2

    #组织要返回的list
    def getResult(self,s,l0):
        obj = Solution()
        l = obj.orderList(l0)
        l2 = list()
        i = 0
        while i < len(s):
            flag = False
            pos = 0
            s0 = ''
            for j in range(len(l)):
                if 2 * j + 1 <= len(l):
                    if i >= l[2 * j] and i <= l[2 * j + 1]:
                        flag = True
                        pos = j
                        break
            if flag:
                s0 = s[l[2 * pos]:l[2 * pos + 1] + 1]
                i = l[2 * pos + 1]
            else:
                s0 = s[i]
            l2.append(s0)
            i += 1
        return l2
