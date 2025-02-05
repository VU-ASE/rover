use std::marker::PhantomData;

pub struct Operating;
pub struct Dormant;

pub struct RoverState<S> {
    pub _ignore: PhantomData<S>,
}
