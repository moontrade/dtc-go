// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE_VLS_SIZE: usize = 22;

/// size                                   u16     = TradeAccountDataUpdateOperationCompleteVLSSize  (22)
/// type                                   u16     = TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE  (10131)
/// base_size                              u16     = TradeAccountDataUpdateOperationCompleteVLSSize  (22)
/// request_id                             u32     = 0
/// is_error                               bool    = false
/// error_text                             string  = ""
/// is_delete_operation                    bool    = false
/// is_symbol_limits_update_operation      bool    = false
/// is_symbol_commission_update_operation  bool    = false
/// trade_account                          string  = ""
pub(crate) const TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE_VLS_DEFAULT: [u8; 22] = [
    22, 0, 147, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait TradeAccountDataUpdateOperationComplete: Message {
    type Safe: TradeAccountDataUpdateOperationComplete;
    type Unsafe: TradeAccountDataUpdateOperationComplete;

    fn request_id(&self) -> u32;

    fn is_error(&self) -> bool;

    fn error_text(&self) -> &str;

    fn is_delete_operation(&self) -> bool;

    fn is_symbol_limits_update_operation(&self) -> bool;

    fn is_symbol_commission_update_operation(&self) -> bool;

    fn trade_account(&self) -> &str;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_is_error(&mut self, value: bool) -> &mut Self;

    fn set_error_text(&mut self, value: &str) -> &mut Self;

    fn set_is_delete_operation(&mut self, value: bool) -> &mut Self;

    fn set_is_symbol_limits_update_operation(&mut self, value: bool) -> &mut Self;

    fn set_is_symbol_commission_update_operation(&mut self, value: bool) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl TradeAccountDataUpdateOperationComplete) {
        to.set_request_id(self.request_id());
        to.set_is_error(self.is_error());
        to.set_error_text(self.error_text());
        to.set_is_delete_operation(self.is_delete_operation());
        to.set_is_symbol_limits_update_operation(self.is_symbol_limits_update_operation());
        to.set_is_symbol_commission_update_operation(self.is_symbol_commission_update_operation());
        to.set_trade_account(self.trade_account());
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

pub struct TradeAccountDataUpdateOperationCompleteVLS {
    data: *const TradeAccountDataUpdateOperationCompleteVLSData,
    capacity: usize,
}

pub struct TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    data: *const TradeAccountDataUpdateOperationCompleteVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct TradeAccountDataUpdateOperationCompleteVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    is_error: bool,
    error_text: VLS,
    is_delete_operation: bool,
    is_symbol_limits_update_operation: bool,
    is_symbol_commission_update_operation: bool,
    trade_account: VLS,
}

impl TradeAccountDataUpdateOperationCompleteVLSData {
    pub fn new() -> Self {
        Self {
            size: 22u16.to_le(),
            r#type: TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE.to_le(),
            base_size: 22u16.to_le(),
            request_id: 0u32.to_le(),
            is_error: false,
            error_text: crate::message::VLS::new(),
            is_delete_operation: false,
            is_symbol_limits_update_operation: false,
            is_symbol_commission_update_operation: false,
            trade_account: crate::message::VLS::new(),
        }
    }
}

unsafe impl Send for TradeAccountDataUpdateOperationCompleteVLS {}
unsafe impl Send for TradeAccountDataUpdateOperationCompleteVLSUnsafe {}
unsafe impl Send for TradeAccountDataUpdateOperationCompleteVLSData {}

impl Drop for TradeAccountDataUpdateOperationCompleteVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountDataUpdateOperationCompleteVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountDataUpdateOperationCompleteVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountDataUpdateOperationCompleteVLS {
    type Target = TradeAccountDataUpdateOperationCompleteVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataUpdateOperationCompleteVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    type Target = TradeAccountDataUpdateOperationCompleteVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for TradeAccountDataUpdateOperationCompleteVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataUpdateOperationCompleteVLS(size: {}, type: {}, base_size: {}, request_id: {}, is_error: {}, error_text: \"{}\", is_delete_operation: {}, is_symbol_limits_update_operation: {}, is_symbol_commission_update_operation: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.is_error(), self.error_text(), self.is_delete_operation(), self.is_symbol_limits_update_operation(), self.is_symbol_commission_update_operation(), self.trade_account()).as_str())
    }
}

impl core::fmt::Debug for TradeAccountDataUpdateOperationCompleteVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataUpdateOperationCompleteVLS(size: {}, type: {}, base_size: {}, request_id: {}, is_error: {}, error_text: \"{}\", is_delete_operation: {}, is_symbol_limits_update_operation: {}, is_symbol_commission_update_operation: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.is_error(), self.error_text(), self.is_delete_operation(), self.is_symbol_limits_update_operation(), self.is_symbol_commission_update_operation(), self.trade_account()).as_str())
    }
}

impl core::fmt::Display for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataUpdateOperationCompleteVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, is_error: {}, error_text: \"{}\", is_delete_operation: {}, is_symbol_limits_update_operation: {}, is_symbol_commission_update_operation: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.is_error(), self.error_text(), self.is_delete_operation(), self.is_symbol_limits_update_operation(), self.is_symbol_commission_update_operation(), self.trade_account()).as_str())
    }
}

impl core::fmt::Debug for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataUpdateOperationCompleteVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, is_error: {}, error_text: \"{}\", is_delete_operation: {}, is_symbol_limits_update_operation: {}, is_symbol_commission_update_operation: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.is_error(), self.error_text(), self.is_delete_operation(), self.is_symbol_limits_update_operation(), self.is_symbol_commission_update_operation(), self.trade_account()).as_str())
    }
}

impl crate::Message for TradeAccountDataUpdateOperationCompleteVLS {
    type Data = TradeAccountDataUpdateOperationCompleteVLSData;

    const TYPE: u16 = TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE;
    const BASE_SIZE: usize = 22;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataUpdateOperationCompleteVLSData::new(),
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
            data: data as *const TradeAccountDataUpdateOperationCompleteVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataUpdateOperationCompleteVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataUpdateOperationCompleteVLSData;
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
impl crate::Message for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    type Data = TradeAccountDataUpdateOperationCompleteVLSData;

    const TYPE: u16 = TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE;
    const BASE_SIZE: usize = 22;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataUpdateOperationCompleteVLSData::new(),
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
            data: data as *const TradeAccountDataUpdateOperationCompleteVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataUpdateOperationCompleteVLSData;
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
impl TradeAccountDataUpdateOperationComplete for TradeAccountDataUpdateOperationCompleteVLS {
    type Safe = TradeAccountDataUpdateOperationCompleteVLS;
    type Unsafe = TradeAccountDataUpdateOperationCompleteVLSUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn is_error(&self) -> bool {
        self.is_error
    }

    fn error_text(&self) -> &str {
        get_vls(self, self.error_text)
    }

    fn is_delete_operation(&self) -> bool {
        self.is_delete_operation
    }

    fn is_symbol_limits_update_operation(&self) -> bool {
        self.is_symbol_limits_update_operation
    }

    fn is_symbol_commission_update_operation(&self) -> bool {
        self.is_symbol_commission_update_operation
    }

    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_is_error(&mut self, value: bool) -> &mut Self {
        self.is_error = value;
        self
    }

    fn set_error_text(&mut self, value: &str) -> &mut Self {
        self.error_text = set_vls(self, self.error_text, value);
        self
    }

    fn set_is_delete_operation(&mut self, value: bool) -> &mut Self {
        self.is_delete_operation = value;
        self
    }

    fn set_is_symbol_limits_update_operation(&mut self, value: bool) -> &mut Self {
        self.is_symbol_limits_update_operation = value;
        self
    }

    fn set_is_symbol_commission_update_operation(&mut self, value: bool) -> &mut Self {
        self.is_symbol_commission_update_operation = value;
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = set_vls(self, self.trade_account, value);
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

impl TradeAccountDataUpdateOperationComplete for TradeAccountDataUpdateOperationCompleteVLSUnsafe {
    type Safe = TradeAccountDataUpdateOperationCompleteVLS;
    type Unsafe = TradeAccountDataUpdateOperationCompleteVLSUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(10) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn is_error(&self) -> bool {
        if self.is_out_of_bounds(11) {
            false
        } else {
            self.is_error
        }
    }

    fn error_text(&self) -> &str {
        if self.is_out_of_bounds(15) {
            ""
        } else {
            get_vls(self, self.error_text)
        }
    }

    fn is_delete_operation(&self) -> bool {
        if self.is_out_of_bounds(16) {
            false
        } else {
            self.is_delete_operation
        }
    }

    fn is_symbol_limits_update_operation(&self) -> bool {
        if self.is_out_of_bounds(17) {
            false
        } else {
            self.is_symbol_limits_update_operation
        }
    }

    fn is_symbol_commission_update_operation(&self) -> bool {
        if self.is_out_of_bounds(18) {
            false
        } else {
            self.is_symbol_commission_update_operation
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(22) {
            ""
        } else {
            get_vls(self, self.trade_account)
        }
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_is_error(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(11) {
            self.is_error = value;
        }
        self
    }

    fn set_error_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(15) {
            self.error_text = set_vls(self, self.error_text, value);
        }
        self
    }

    fn set_is_delete_operation(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.is_delete_operation = value;
        }
        self
    }

    fn set_is_symbol_limits_update_operation(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(17) {
            self.is_symbol_limits_update_operation = value;
        }
        self
    }

    fn set_is_symbol_commission_update_operation(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(18) {
            self.is_symbol_commission_update_operation = value;
        }
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(22) {
            self.trade_account = set_vls(self, self.trade_account, value);
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
                22usize,
                core::mem::size_of::<TradeAccountDataUpdateOperationCompleteVLSData>(),
                "TradeAccountDataUpdateOperationCompleteVLSData sizeof expected {:} but was {:}",
                22usize,
                core::mem::size_of::<TradeAccountDataUpdateOperationCompleteVLSData>()
            );
            assert_eq!(
                22u16,
                TradeAccountDataUpdateOperationCompleteVLS::new().size(),
                "TradeAccountDataUpdateOperationCompleteVLS sizeof expected {:} but was {:}",
                22u16,
                TradeAccountDataUpdateOperationCompleteVLS::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE,
                TradeAccountDataUpdateOperationCompleteVLS::new().r#type(),
                "TradeAccountDataUpdateOperationCompleteVLS type expected {:} but was {:}",
                TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE,
                TradeAccountDataUpdateOperationCompleteVLS::new().r#type(),
            );
            assert_eq!(
                10131u16,
                TradeAccountDataUpdateOperationCompleteVLS::new().r#type(),
                "TradeAccountDataUpdateOperationCompleteVLS type expected {:} but was {:}",
                10131u16,
                TradeAccountDataUpdateOperationCompleteVLS::new().r#type(),
            );
            let d = TradeAccountDataUpdateOperationCompleteVLSData::new();
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
                (core::ptr::addr_of!(d.is_error) as usize) - p,
                "is_error offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.is_error) as usize) - p,
            );
            assert_eq!(
                11usize,
                (core::ptr::addr_of!(d.error_text) as usize) - p,
                "error_text offset expected {:} but was {:}",
                11usize,
                (core::ptr::addr_of!(d.error_text) as usize) - p,
            );
            assert_eq!(
                15usize,
                (core::ptr::addr_of!(d.is_delete_operation) as usize) - p,
                "is_delete_operation offset expected {:} but was {:}",
                15usize,
                (core::ptr::addr_of!(d.is_delete_operation) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.is_symbol_limits_update_operation) as usize) - p,
                "is_symbol_limits_update_operation offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.is_symbol_limits_update_operation) as usize) - p,
            );
            assert_eq!(
                17usize,
                (core::ptr::addr_of!(d.is_symbol_commission_update_operation) as usize) - p,
                "is_symbol_commission_update_operation offset expected {:} but was {:}",
                17usize,
                (core::ptr::addr_of!(d.is_symbol_commission_update_operation) as usize) - p,
            );
            assert_eq!(
                18usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                18usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
        }
    }
}
