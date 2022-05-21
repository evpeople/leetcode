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

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
struct Solution;
use std::borrow::Borrow;
use std::cell::RefCell;
use std::collections::HashMap;
use std::rc::Rc;
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
    pub fn longest_palindrome(s: String) -> String {
        let (mut start, mut end) = (0, 0);
        let s = s.chars().collect::<Vec<_>>();

        let mut dp = vec![vec![false; s.len()]; s.len()];
        for i in (0..s.len()).rev() {
            for j in i..s.len() {
                if i == j || j - i == 1 && s[i] == s[j] {
                    dp[i][j] = true;
                } else {
                    dp[i][j] = dp[i + 1][j - 1] && s[i] == s[j];
                }
                if dp[i][j] && j - i > end - start {
                    start = i;
                    end = j;
                }
            }
        }

        s[start..=end].iter().collect()
    }
}

use futures::executor::block_on;

struct Song {
    author: String,
    name: String,
}

async fn learn_song() -> Song {
    Song {
        author: "曲婉婷".to_string(),
        name: String::from("《我的歌声里》"),
    }
}

async fn sing_song(song: Song) {
    println!(
        "给大家献上一首{}的{} ~ {}",
        song.author, song.name, "你存在我深深的脑海里~ ~"
    );
}

async fn dance() {
    println!("唱到情深处，身体不由自主的动了起来~ ~");
}

async fn learn_and_sing() {
    // 这里使用`.await`来等待学歌的完成，但是并不会阻塞当前线程，该线程在学歌的任务`.await`后，完全可以去执行跳舞的任务
    let song = learn_song().await;

    // 唱歌必须要在学歌之后
    sing_song(song).await;
}

async fn async_main() {
    let f1 = learn_and_sing();
    let f2 = dance();

    // `join!`可以并发的处理和等待多个`Future`，若`learn_and_sing Future`被阻塞，那`dance Future`可以拿过线程的所有权继续执行。若`dance`也变成阻塞状态，那`learn_and_sing`又可以再次拿回线程所有权，继续执行。
    // 若两个都被阻塞，那么`async main`会变成阻塞状态，然后让出线程所有权，并将其交给`main`函数中的`block_on`执行器
    futures::join!(f2, f1);
}

fn main() {
    block_on(async_main());
}

// fn main() {
//     println!("Hello, world!");
//     let a = Solution::longest_palindrome(String::from("asdsa"));
// }
