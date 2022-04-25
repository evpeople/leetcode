#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}
struct Solution {}
use std::collections::HashMap;
impl Solution {
    pub fn first_uniq_char(s: String) -> char {
        let mut map = HashMap::new();
        // let s = s.chars().collect()::<Vec<_>>();
        let s = s.chars().collect::<Vec<_>>();
        s.iter().for_each(|c| *map.entry(c).or_insert(0) += 1);
        if let Some(&c) = s.iter().find(|&c| map[c] == 1) {
            c
        } else {
            ' '
        }
    }
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list1, list2) {
            (None, None) => None,
            (None, r) => r,
            (l, None) => l,
            (Some(mut l), Some(mut r)) => {
                if l.val <= r.val {
                    l.next = Self::merge_two_lists(l.next, Some(r));
                    Some(l)
                } else {
                    r.next = Self::merge_two_lists(Some(l), r.next);
                    Some(r)
                }
            }
        }
    }
    pub fn add_strings<S: AsRef<str>>(num1: S, num2: S) -> String {
        let (mut num1, mut num2) = (num1.as_ref().bytes().rev(), num2.as_ref().bytes().rev());
        let mut ret = vec![];
        let mut carry = 0;
        loop {
            match (num1.next(), num2.next()) {
                (None, None) => break,
                (a, b) => {
                    let mut n = a.unwrap_or(b'0') + b.unwrap_or(b'0') - b'0' + carry;
                    carry = 0;
                    if n > b'9' {
                        n -= 10;
                        carry = 1;
                    }
                    ret.push(n as char);
                }
            }
        }
        if carry != 0 {
            ret.push('1');
        }
        ret.iter().rev().collect()
    }
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        if matrix.is_empty() || matrix[0].is_empty() {
            return vec![];
        }
        let mut res = Vec::new();
        let (mut row_begin, mut row_end) = (0, matrix.len() - 1);
        let (mut col_begin, mut col_end) = (0, matrix[0].len() - 1);

        while row_begin <= row_end && col_begin <= col_end {
            (col_begin..=col_end).for_each(|i| res.push(matrix[row_begin][i]));
            (row_begin + 1..row_end).for_each(|i| res.push(matrix[i][col_end]));

            if row_begin != row_end {
                (col_begin..=col_end)
                    .rev()
                    .for_each(|i| res.push(matrix[row_end][i]));
            }
            if col_begin != col_end {
                (row_begin + 1..row_end)
                    .rev()
                    .for_each(|i| res.push(matrix[i][col_begin]))
            }

            row_begin += 1;
            row_end = row_end.saturating_sub(1);
            col_begin += 1;
            col_end = col_end.saturating_sub(1);
        }
        res
    }
}

fn main() {
    println!("Hello, world!");
}
