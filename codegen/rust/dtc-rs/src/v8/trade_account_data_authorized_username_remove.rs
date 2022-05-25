// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE_VLS_SIZE: usize = 18;

/// size           u16     = TradeAccountDataAuthorizedUsernameRemoveVLSSize  (18)
/// type           u16     = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE  (10126)
/// base_size      u16     = TradeAccountDataAuthorizedUsernameRemoveVLSSize  (18)
/// request_id     u32     = 0
/// trade_account  string  = ""
/// username       string  = ""
pub(crate) const TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE_VLS_DEFAULT: [u8; 18] =
    [18, 0, 142, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait TradeAccountDataAuthorizedUsernameRemove: Message {
    type Safe: TradeAccountDataAuthorizedUsernameRemove;
    type Unsafe: TradeAccountDataAuthorizedUsernameRemove;

    fn request_id(&self) -> u32;

    fn trade_account(&self) -> &str;

    fn username(&self) -> &str;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn set_username(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl TradeAccountDataAuthorizedUsernameRemove) {
        to.set_request_id(self.request_id());
        to.set_trade_account(self.trade_account());
        to.set_username(self.username());
    }
}

pub struct TradeAccountDataAuthorizedUsernameRemoveVLS {
    data: *const TradeAccountDataAuthorizedUsernameRemoveVLSData,
    capacity: usize,
}

pub struct TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {
    data: *const TradeAccountDataAuthorizedUsernameRemoveVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct TradeAccountDataAuthorizedUsernameRemoveVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    trade_account: VLS,
    username: VLS,
}

impl TradeAccountDataAuthorizedUsernameRemoveVLSData {
    pub fn new() -> Self {
        Self {
            size: 18u16.to_le(),
            r#type: TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE.to_le(),
            base_size: 18u16.to_le(),
            request_id: 0u32.to_le(),
            trade_account: crate::message::VLS::new(),
            username: crate::message::VLS::new(),
        }
    }
}

unsafe impl Send for TradeAccountDataAuthorizedUsernameRemoveVLS {}
unsafe impl Send for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {}
unsafe impl Send for TradeAccountDataAuthorizedUsernameRemoveVLSData {}

impl Drop for TradeAccountDataAuthorizedUsernameRemoveVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountDataAuthorizedUsernameRemoveVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountDataAuthorizedUsernameRemoveVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountDataAuthorizedUsernameRemoveVLS {
    type Target = TradeAccountDataAuthorizedUsernameRemoveVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataAuthorizedUsernameRemoveVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {
    type Target = TradeAccountDataAuthorizedUsernameRemoveVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for TradeAccountDataAuthorizedUsernameRemoveVLS {
    type Data = TradeAccountDataAuthorizedUsernameRemoveVLSData;

    const BASE_SIZE: usize = 18;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataAuthorizedUsernameRemoveVLSData::new(),
            ),
            capacity: Self::DEFAULT_CAPACITY,
        }
    }

    #[inline]
    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }

    #[inline]
    fn base_size(&self) -> u16 {
        u16::from_le(self.base_size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        self.capacity as u16
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {
        Self {
            data: data as *const TradeAccountDataAuthorizedUsernameRemoveVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataAuthorizedUsernameRemoveVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataAuthorizedUsernameRemoveVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }
}
impl crate::Message for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {
    type Data = TradeAccountDataAuthorizedUsernameRemoveVLSData;

    const BASE_SIZE: usize = 18;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataAuthorizedUsernameRemoveVLSData::new(),
            ),
            capacity: Self::DEFAULT_CAPACITY,
        }
    }

    #[inline]
    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }

    #[inline]
    fn base_size(&self) -> u16 {
        u16::from_le(self.base_size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        self.capacity as u16
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {
        Self {
            data: data as *const TradeAccountDataAuthorizedUsernameRemoveVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataAuthorizedUsernameRemoveVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }
}
impl TradeAccountDataAuthorizedUsernameRemove for TradeAccountDataAuthorizedUsernameRemoveVLS {
    type Safe = TradeAccountDataAuthorizedUsernameRemoveVLS;
    type Unsafe = TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    fn username(&self) -> &str {
        get_vls(self, self.username)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = set_vls(self, self.trade_account, value);
        self
    }

    fn set_username(&mut self, value: &str) -> &mut Self {
        self.username = set_vls(self, self.username, value);
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
    }
}

impl TradeAccountDataAuthorizedUsernameRemove
    for TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe
{
    type Safe = TradeAccountDataAuthorizedUsernameRemoveVLS;
    type Unsafe = TradeAccountDataAuthorizedUsernameRemoveVLSUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(10) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(14) {
            ""
        } else {
            get_vls(self, self.trade_account)
        }
    }

    fn username(&self) -> &str {
        if self.is_out_of_bounds(18) {
            ""
        } else {
            get_vls(self, self.username)
        }
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(14) {
            self.trade_account = set_vls(self, self.trade_account, value);
        }
        self
    }

    fn set_username(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(18) {
            self.username = set_vls(self, self.username, value);
        }
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }
}

#[cfg(test)]
pub(crate) mod tests {
    use super::*;

    #[test]
    pub(crate) fn layout() {
        unsafe {
            assert_eq!(
                18usize,
                core::mem::size_of::<TradeAccountDataAuthorizedUsernameRemoveVLSData>(),
                "TradeAccountDataAuthorizedUsernameRemoveVLSData sizeof expected {:} but was {:}",
                18usize,
                core::mem::size_of::<TradeAccountDataAuthorizedUsernameRemoveVLSData>()
            );
            assert_eq!(
                18u16,
                TradeAccountDataAuthorizedUsernameRemoveVLS::new().size(),
                "TradeAccountDataAuthorizedUsernameRemoveVLS sizeof expected {:} but was {:}",
                18u16,
                TradeAccountDataAuthorizedUsernameRemoveVLS::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE,
                TradeAccountDataAuthorizedUsernameRemoveVLS::new().r#type(),
                "TradeAccountDataAuthorizedUsernameRemoveVLS type expected {:} but was {:}",
                TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE,
                TradeAccountDataAuthorizedUsernameRemoveVLS::new().r#type(),
            );
            assert_eq!(
                10126u16,
                TradeAccountDataAuthorizedUsernameRemoveVLS::new().r#type(),
                "TradeAccountDataAuthorizedUsernameRemoveVLS type expected {:} but was {:}",
                10126u16,
                TradeAccountDataAuthorizedUsernameRemoveVLS::new().r#type(),
            );
            let d = TradeAccountDataAuthorizedUsernameRemoveVLSData::new();
            let p = (&d as *const _ as *const u8).offset(0) as usize;
            assert_eq!(
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
                "size offset expected {:} but was {:}",
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
            );
            assert_eq!(
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
                "type offset expected {:} but was {:}",
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
            );
            assert_eq!(
                4usize,
                (core::ptr::addr_of!(d.base_size) as usize) - p,
                "base_size offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.base_size) as usize) - p,
            );
            assert_eq!(
                6usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
            );
            assert_eq!(
                10usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
            assert_eq!(
                14usize,
                (core::ptr::addr_of!(d.username) as usize) - p,
                "username offset expected {:} but was {:}",
                14usize,
                (core::ptr::addr_of!(d.username) as usize) - p,
            );
        }
    }
}