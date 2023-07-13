Для заданного целочисленного массива numsвозвращайте значение, trueесли какое-либо значение встречается в массиве не менее двух разfalse , и возвращайте значение , если все элементы различны.

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true


Попробуйте написать код самостоятельно, перечислив подход 🥺.
Brute Force : идея кода

Сделать вложенный цикл, сгенерировать все возможные пары
Поставьте условие, если оба числа, сгенерированные в паре, одинаковы.
В этом подходе будет сформирована только уникальная пара, потому что внешний цикл работает от 0 до n - 1, а внутренний цикл начнется с одного значения, дополнительного к предыдущему значению цикла (которое заставить его запустить n*(n+1)/2 ) . если мы сопоставляем каждую пару векторов, то, возможно, мы можем сравнить, имеют ли какие-либо из них одинаковое значение, а затем вернуть true. иначе в конце цикла вложенной формы вернуть false.
// Brute Force
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        bool flag = false;
        for(int i =0;i<nums.size();i++){
            for(int j=i+1;j<nums.size();j++){
                if(nums[i] == nums[j]) return true;
            }
        }
        return flag;
    }
};
Сложность
Временная сложность:O((n∗(n+1))/2) ≈ O(n^2)
Сложность пространства:О(1)
Укороченный подход : идея кода

отсортировать массив.
линейно пройти, найти, если есть какое-либо число и его большее число равны или нет
// Shorted Approach
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        sort(nums.begin(),nums.end());
        bool flag = false;
        for(int i =0;i<nums.size()-1;i++){
            if(nums[i] == nums[i+1]) return true;
        }
        return flag;
    }
};
Сложность
Временная сложность:O(n∗log(n))
Сложность пространства:O(1)
Установить подход : идея, лежащая в основе кода

Набор только не содержит повторяющихся элементов.

Если размер набора меньше исходного вектора, то он содержит повторяющийся элемент.

Для вашего подхода к набору временная сложность может быть уменьшена с O (nlogn) до O (n), используя unordered_set. Обычный набор использует деревья со вставкой O(logn), но unordered_sets использует хеширование для O(1). :)

// Set Approach
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        return nums.size() > set<int>(nums.begin(),nums.end()).size();
    }
};
Сложность
Временная сложность:O(n∗log(n))
Сложность пространства: O(n)

Карта : Идея, лежащая в основе кода
вектор обхода и частота подсчета с использованием любой структуры данных.
Переберите структуру данных, чтобы определить, превышает ли частота значение 1.
// Contains Duplicate
  class Solution {
  public:
      bool containsDuplicate(vector<int>& nums) {
          map<int,int> mp;
          for(auto i : nums) mp[i]++;
          bool flag = false;
          for(auto i : mp){
              if(i.second >= 2) return true;
          }
          return flag;
      }
  };
Сложность
Time complexity: O(n∗log(n))
Space complexity: O(n)


Hashmap : идея кода
вектор обхода и частота подсчета с использованием любой структуры данных.
Переберите структуру данных, чтобы определить, превышает ли частота значение 1.
то же, что и выше, но разница в том, что это будет вO(n)
// Contains Duplicate
  class Solution {
  public:
      bool containsDuplicate(vector<int>& nums) {
          unordered_map<int,int> mp;
          for(auto i : nums) mp[i]++;
          bool flag = false;
          for(auto i : mp){
              if(i.second >= 2) return true;
          }
          return flag;
      }
  };
Сложность
Time complexity: O(n)
Space complexity: O(n)