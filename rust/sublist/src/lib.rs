#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(first_list: &[T], second_list: &[T]) -> Comparison {
    use Comparison::*;
    match (first_list.len(), second_list.len()) {
        (0, 0) => Equal,
        (0, _) => Sublist,
        (_, 0) => Superlist,
        (a, b) if a > b && first_list.windows(b).any(|v| v == second_list) => Superlist,
        (a, b) if a < b && second_list.windows(a).any(|v| v == first_list) => Sublist,
        (_, _) if first_list == second_list => Equal,
        (_, _) => Unequal,
    }
}
