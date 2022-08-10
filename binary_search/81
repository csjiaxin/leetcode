from typing import List
class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        l=0 
        r=len(nums)-1 

        while(l <= r):
            mid = l +(r-l)//2
            if nums[mid] == target:
                return True
            elif nums[mid] < nums[r]:
                if target > nums[mid] and target <= nums[r]:
                    l=mid+1
                else:
                    r=mid-1
            else:
            # elif nums[mid] > nums[l]:
                if target >=nums[l] and target < nums[mid]:
                    r=mid-1
                else:
                    l=mid+1

        return False

s=Solution()
print(s.search([1,0,1,1,1], 0))
print(s.search([1,3], 3))
print(s.search([4,5,6,7,0,1,2], 3))
print(s.search([4,5,6,7,8,1,2,3], 8))
