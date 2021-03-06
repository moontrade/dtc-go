// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD_VLS_SIZE: usize = 19;

/// size           u16     = TradeAccountDataUsernameToShareWithAddVLSSize  (19)
/// type           u16     = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD  (10128)
/// base_size      u16     = TradeAccountDataUsernameToShareWithAddVLSSize  (19)
/// request_id     u32     = 0
/// trade_account  string  = ""
/// username       string  = ""
/// is_read_only   bool    = false
pub(crate) const TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD_VLS_DEFAULT: [u8; 19] =
    [19, 0, 144, 39, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait TradeAccountDataUsernameToShareWithAdd: Message {
    type Safe: TradeAccountDataUsernameToShareWithAdd;
    type Unsafe: TradeAccountDataUsernameToShareWithAdd;

    fn request_id(&self) -> u32;

    fn trade_account(&self) -> &str;

    fn username(&self) -> &str;

    fn is_read_only(&self) -> bool;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn set_username(&mut self, value: &str) -> &mut Self;

    fn set_is_read_only(&mut self, value: bool) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl TradeAccountDataUsernameToShareWithAdd) {
        to.set_request_id(self.request_id());
        to.set_trade_account(self.trade_account());
        to.set_username(self.username());
        to.set_is_read_only(self.is_read_only());
    }

    #[inline]
    fn parse<F: Fn(Parsed<Self::Safe, Self::Unsafe>) -> Result<(), crate::Error>>(
        data: &[u8],
        f: F,
    ) -> Result<(), crate::Error> {
        if data.len() < 6 {
            return Err(crate::Error::Malformed("need more data"));
        }
        let size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) };
        let base_size = if Self::BASE_SIZE_OFFSET == 0 {
            size
        } else {
            let base_size = unsafe {
                u16::from_le(*(data.as_ptr().offset(Self::BASE_SIZE_OFFSET) as *const u16))
            };
            if base_size > size {
                return Err(crate::Error::Malformed("base_size is greater than size"));
            }
            base_size
        };
        if (base_size as usize) >= Self::BASE_SIZE {
            let msg = unsafe { Self::Safe::from_raw_parts(data.as_ptr(), size as usize) };
            let r = f(Parsed::Safe(&msg));
            core::mem::forget(msg);
            r
        } else {
            let msg = unsafe { Self::Unsafe::from_raw_parts(data.as_ptr(), size as usize) };
            let r = f(Parsed::Unsafe(&msg));
            core::mem::forget(msg);
            r
        }
    }
}

pub struct TradeAccountDataUsernameToShareWithAddVLS {
    data: *const TradeAccountDataUsernameToShareWithAddVLSData,
    capacity: usize,
}

pub struct TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    data: *const TradeAccountDataUsernameToShareWithAddVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct TradeAccountDataUsernameToShareWithAddVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    trade_account: VLS,
    username: VLS,
    is_read_only: bool,
}

impl TradeAccountDataUsernameToShareWithAddVLSData {
    pub fn new() -> Self {
        Self {
            size: 19u16.to_le(),
            r#type: TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD.to_le(),
            base_size: 19u16.to_le(),
            request_id: 0u32.to_le(),
            trade_account: crate::message::VLS::new(),
            username: crate::message::VLS::new(),
            is_read_only: false,
        }
    }
}

unsafe impl Send for TradeAccountDataUsernameToShareWithAddVLS {}
unsafe impl Send for TradeAccountDataUsernameToShareWithAddVLSUnsafe {}
unsafe impl Send for TradeAccountDataUsernameToShareWithAddVLSData {}

impl Drop for TradeAccountDataUsernameToShareWithAddVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountDataUsernameToShareWithAddVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountDataUsernameToShareWithAddVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountDataUsernameToShareWithAddVLS {
    type Target = TradeAccountDataUsernameToShareWithAddVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataUsernameToShareWithAddVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    type Target = TradeAccountDataUsernameToShareWithAddVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for TradeAccountDataUsernameToShareWithAddVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataUsernameToShareWithAddVLS(size: {}, type: {}, base_size: {}, request_id: {}, trade_account: \"{}\", username: \"{}\", is_read_only: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.trade_account(), self.username(), self.is_read_only()).as_str())
    }
}

impl core::fmt::Debug for TradeAccountDataUsernameToShareWithAddVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataUsernameToShareWithAddVLS(size: {}, type: {}, base_size: {}, request_id: {}, trade_account: \"{}\", username: \"{}\", is_read_only: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.trade_account(), self.username(), self.is_read_only()).as_str())
    }
}

impl core::fmt::Display for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataUsernameToShareWithAddVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, trade_account: \"{}\", username: \"{}\", is_read_only: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.trade_account(), self.username(), self.is_read_only()).as_str())
    }
}

impl core::fmt::Debug for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataUsernameToShareWithAddVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, trade_account: \"{}\", username: \"{}\", is_read_only: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.trade_account(), self.username(), self.is_read_only()).as_str())
    }
}

impl crate::Message for TradeAccountDataUsernameToShareWithAddVLS {
    type Data = TradeAccountDataUsernameToShareWithAddVLSData;

    const TYPE: u16 = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD;
    const BASE_SIZE: usize = 19;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataUsernameToShareWithAddVLSData::new(),
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
            data: data as *const TradeAccountDataUsernameToShareWithAddVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataUsernameToShareWithAddVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataUsernameToShareWithAddVLSData;
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
impl crate::Message for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    type Data = TradeAccountDataUsernameToShareWithAddVLSData;

    const TYPE: u16 = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD;
    const BASE_SIZE: usize = 19;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataUsernameToShareWithAddVLSData::new(),
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
            data: data as *const TradeAccountDataUsernameToShareWithAddVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataUsernameToShareWithAddVLSData;
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
impl TradeAccountDataUsernameToShareWithAdd for TradeAccountDataUsernameToShareWithAddVLS {
    type Safe = TradeAccountDataUsernameToShareWithAddVLS;
    type Unsafe = TradeAccountDataUsernameToShareWithAddVLSUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    fn username(&self) -> &str {
        get_vls(self, self.username)
    }

    fn is_read_only(&self) -> bool {
        self.is_read_only
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

    fn set_is_read_only(&mut self, value: bool) -> &mut Self {
        self.is_read_only = value;
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

impl TradeAccountDataUsernameToShareWithAdd for TradeAccountDataUsernameToShareWithAddVLSUnsafe {
    type Safe = TradeAccountDataUsernameToShareWithAddVLS;
    type Unsafe = TradeAccountDataUsernameToShareWithAddVLSUnsafe;

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

    fn is_read_only(&self) -> bool {
        if self.is_out_of_bounds(19) {
            false
        } else {
            self.is_read_only
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

    fn set_is_read_only(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(19) {
            self.is_read_only = value;
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
                19usize,
                core::mem::size_of::<TradeAccountDataUsernameToShareWithAddVLSData>(),
                "TradeAccountDataUsernameToShareWithAddVLSData sizeof expected {:} but was {:}",
                19usize,
                core::mem::size_of::<TradeAccountDataUsernameToShareWithAddVLSData>()
            );
            assert_eq!(
                19u16,
                TradeAccountDataUsernameToShareWithAddVLS::new().size(),
                "TradeAccountDataUsernameToShareWithAddVLS sizeof expected {:} but was {:}",
                19u16,
                TradeAccountDataUsernameToShareWithAddVLS::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD,
                TradeAccountDataUsernameToShareWithAddVLS::new().r#type(),
                "TradeAccountDataUsernameToShareWithAddVLS type expected {:} but was {:}",
                TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD,
                TradeAccountDataUsernameToShareWithAddVLS::new().r#type(),
            );
            assert_eq!(
                10128u16,
                TradeAccountDataUsernameToShareWithAddVLS::new().r#type(),
                "TradeAccountDataUsernameToShareWithAddVLS type expected {:} but was {:}",
                10128u16,
                TradeAccountDataUsernameToShareWithAddVLS::new().r#type(),
            );
            let d = TradeAccountDataUsernameToShareWithAddVLSData::new();
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
            assert_eq!(
                18usize,
                (core::ptr::addr_of!(d.is_read_only) as usize) - p,
                "is_read_only offset expected {:} but was {:}",
                18usize,
                (core::ptr::addr_of!(d.is_read_only) as usize) - p,
            );
        }
    }
}
