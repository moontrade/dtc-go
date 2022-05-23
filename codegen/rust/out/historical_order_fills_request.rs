// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-23 21:07:20.84591 +0800 WITA m=+0.009860584
use super::constants::*;
use super::alias::*;
use super::enums::*;
use crate::message::*;

const HISTORICAL_ORDER_FILLS_REQUEST_VLS_SIZE: usize = 32;

const HISTORICAL_ORDER_FILLS_REQUEST_FIXED_SIZE: usize = 88;

/// size             u16       = HistoricalOrderFillsRequestVLSSize  (32)
/// r#type           u16       = HISTORICAL_ORDER_FILLS_REQUEST  (303)
/// base_size        u16       = HistoricalOrderFillsRequestVLSSize  (32)
/// request_id       i32       = 0
/// server_order_id  string    = ""
/// number_of_days   i32       = 0
/// trade_account    string    = ""
/// start_date_time  DateTime  = 0
const HISTORICAL_ORDER_FILLS_REQUEST_VLS_DEFAULT: [u8; 32] = [32, 0, 47, 1, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size             u16       = HistoricalOrderFillsRequestFixedSize  (88)
/// r#type           u16       = HISTORICAL_ORDER_FILLS_REQUEST  (303)
/// request_id       i32       = 0
/// server_order_id  string32  = ""
/// number_of_days   i32       = 0
/// trade_account    string32  = ""
/// start_date_time  DateTime  = 0
const HISTORICAL_ORDER_FILLS_REQUEST_FIXED_DEFAULT: [u8; 88] = [88, 0, 47, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// This is a message from the Client to the Server to request order fills/executions
/// for an order or orders.
pub trait HistoricalOrderFillsRequest {
    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn request_id(&self) -> i32;

    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn server_order_id(&self) -> &str;

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn number_of_days(&self) -> i32;

    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn trade_account(&self) -> &str;

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn start_date_time(&self) -> DateTime;

    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn set_server_order_id(&mut self, value: &str) -> &mut Self;

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_number_of_days(&mut self, value: i32) -> &mut Self;

    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self;

    fn copy_to(&self, to: &mut impl HistoricalOrderFillsRequest) {
        to.set_request_id(self.request_id());
        to.set_server_order_id(self.server_order_id());
        to.set_number_of_days(self.number_of_days());
        to.set_trade_account(self.trade_account());
        to.set_start_date_time(self.start_date_time());
    }
}

/// This is a message from the Client to the Server to request order fills/executions
/// for an order or orders.
pub struct HistoricalOrderFillsRequestVLS {
    data: *const HistoricalOrderFillsRequestVLSData,
    capacity: usize
}

pub struct HistoricalOrderFillsRequestVLSUnsafe {
    data: *const HistoricalOrderFillsRequestVLSData,
    capacity: usize
}

#[repr(packed, C)]
pub struct HistoricalOrderFillsRequestVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    server_order_id: VLS,
    number_of_days: i32,
    trade_account: VLS,
    start_date_time: DateTime,
}

/// This is a message from the Client to the Server to request order fills/executions
/// for an order or orders.
pub struct HistoricalOrderFillsRequestFixed {
    data: *const HistoricalOrderFillsRequestFixedData
}

pub struct HistoricalOrderFillsRequestFixedUnsafe {
    data: *const HistoricalOrderFillsRequestFixedData
}

#[repr(packed, C)]
pub struct HistoricalOrderFillsRequestFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    server_order_id: [u8; 32],
    number_of_days: i32,
    trade_account: [u8; 32],
    start_date_time: DateTime,
}

impl HistoricalOrderFillsRequestVLSData {
    pub fn new() -> Self {
        Self {
            size: 32u16.to_le(),
            r#type: HISTORICAL_ORDER_FILLS_REQUEST.to_le(),
            base_size: 32u16.to_le(),
            request_id: 0,
            server_order_id: crate::message::VLS::new(),
            number_of_days: 0,
            trade_account: crate::message::VLS::new(),
            start_date_time: 0,
        }
    }
}

impl HistoricalOrderFillsRequestFixedData {
    pub fn new() -> Self {
        Self {
            size: 88u16.to_le(),
            r#type: HISTORICAL_ORDER_FILLS_REQUEST.to_le(),
            request_id: 0,
            server_order_id: [0; 32],
            number_of_days: 0,
            trade_account: [0; 32],
            start_date_time: 0,
        }
    }
}

unsafe impl Send for HistoricalOrderFillsRequestFixed {}
unsafe impl Send for HistoricalOrderFillsRequestFixedUnsafe {}
unsafe impl Send for HistoricalOrderFillsRequestFixedData {}
unsafe impl Send for HistoricalOrderFillsRequestVLS {}
unsafe impl Send for HistoricalOrderFillsRequestVLSUnsafe {}
unsafe impl Send for HistoricalOrderFillsRequestVLSData {}

impl Drop for HistoricalOrderFillsRequestFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalOrderFillsRequestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalOrderFillsRequestVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalOrderFillsRequestVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for HistoricalOrderFillsRequestFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalOrderFillsRequestFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalOrderFillsRequestVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalOrderFillsRequestVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for HistoricalOrderFillsRequestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalOrderFillsRequestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalOrderFillsRequestVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalOrderFillsRequestVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for HistoricalOrderFillsRequestFixed {
    type Target = HistoricalOrderFillsRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalOrderFillsRequestFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalOrderFillsRequestFixedUnsafe {
    type Target = HistoricalOrderFillsRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalOrderFillsRequestFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalOrderFillsRequestVLS {
    type Target = HistoricalOrderFillsRequestVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalOrderFillsRequestVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalOrderFillsRequestVLSUnsafe {
    type Target = HistoricalOrderFillsRequestVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalOrderFillsRequestVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for HistoricalOrderFillsRequestFixed {
    type Safe = HistoricalOrderFillsRequestFixed;
    type Unsafe = HistoricalOrderFillsRequestFixedUnsafe;
    type Data = HistoricalOrderFillsRequestFixedData;
    const BASE_SIZE: usize = 88;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalOrderFillsRequestFixedData::new())
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
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
        u16::from_le(self.size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {
        Self {
            data: data as *const HistoricalOrderFillsRequestFixedData
        }
    }

}

impl crate::Message for HistoricalOrderFillsRequestFixedUnsafe {
    type Safe = HistoricalOrderFillsRequestFixed;
    type Unsafe = HistoricalOrderFillsRequestFixedUnsafe;
    type Data = HistoricalOrderFillsRequestFixedData;
    const BASE_SIZE: usize = 88;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalOrderFillsRequestFixedData::new())
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
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
        u16::from_le(self.size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {
        Self {
            data: data as *const HistoricalOrderFillsRequestFixedData
        }
    }

}

impl crate::Message for HistoricalOrderFillsRequestVLS {
    type Safe = HistoricalOrderFillsRequestVLS;
    type Unsafe = HistoricalOrderFillsRequestVLSUnsafe;
    type Data = HistoricalOrderFillsRequestVLSData;
    const BASE_SIZE: usize = 32;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalOrderFillsRequestVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
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
            data: data as *const HistoricalOrderFillsRequestVLSData,
            capacity
        }
    }

}

impl crate::VLSMessage for HistoricalOrderFillsRequestVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalOrderFillsRequestVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }

}impl crate::Message for HistoricalOrderFillsRequestVLSUnsafe {
    type Safe = HistoricalOrderFillsRequestVLS;
    type Unsafe = HistoricalOrderFillsRequestVLSUnsafe;
    type Data = HistoricalOrderFillsRequestVLSData;
    const BASE_SIZE: usize = 32;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalOrderFillsRequestVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
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
            data: data as *const HistoricalOrderFillsRequestVLSData,
            capacity
        }
    }

}

impl crate::VLSMessage for HistoricalOrderFillsRequestVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalOrderFillsRequestVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }

}/// This is a message from the Client to the Server to request order fills/executions
/// for an order or orders.
impl HistoricalOrderFillsRequest for HistoricalOrderFillsRequestVLS {
    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn server_order_id(&self) -> &str {
        crate::get_vls(self, self.server_order_id)
    }

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn number_of_days(&self) -> i32 {
        i32::from_le(self.number_of_days)
    }

    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn trade_account(&self) -> &str {
        crate::get_vls(self, self.trade_account)
    }

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn start_date_time(&self) -> DateTime {
        i64::from_le(self.start_date_time)
    }

    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }


    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn set_server_order_id(&mut self, value: &str) -> &mut Self {
        self.server_order_id = crate::set_vls(self, self.server_order_id, value);
        self
    }


    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_number_of_days(&mut self, value: i32) -> &mut Self {
        self.number_of_days = value.to_le();
        self
    }


    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = crate::set_vls(self, self.trade_account, value);
        self
    }


    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        self.start_date_time = value.to_le();
        self
    }

}

/// This is a message from the Client to the Server to request order fills/executions
/// for an order or orders.
impl HistoricalOrderFillsRequest for HistoricalOrderFillsRequestVLSUnsafe {
    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn server_order_id(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            crate::get_vls(self, self.server_order_id)
        }
    }

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn number_of_days(&self) -> i32 {
        if self.is_out_of_bounds(20) {
            0
        } else {
            i32::from_le(self.number_of_days)
        }
    }

    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(24) {
            ""
        } else {
            crate::get_vls(self, self.trade_account)
        }
    }

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn start_date_time(&self) -> DateTime {
        if self.is_out_of_bounds(32) {
            0
        } else {
            i64::from_le(self.start_date_time)
        }
    }

    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_id = value.to_le();
        }
        self
    }


    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn set_server_order_id(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.server_order_id = crate::set_vls(self, self.server_order_id, value);
        }
        self
    }


    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_number_of_days(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.number_of_days = value.to_le();
        }
        self
    }


    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.trade_account = crate::set_vls(self, self.trade_account, value);
        }
        self
    }


    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        if !self.is_out_of_bounds(32) {
            self.start_date_time = value.to_le();
        }
        self
    }

}

/// This is a message from the Client to the Server to request order fills/executions
/// for an order or orders.
impl HistoricalOrderFillsRequest for HistoricalOrderFillsRequestFixed {
    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn server_order_id(&self) -> &str {
        crate::get_fixed(&self.server_order_id[..])
    }

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn number_of_days(&self) -> i32 {
        i32::from_le(self.number_of_days)
    }

    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn trade_account(&self) -> &str {
        crate::get_fixed(&self.trade_account[..])
    }

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn start_date_time(&self) -> DateTime {
        i64::from_le(self.start_date_time)
    }

    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }


    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn set_server_order_id(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.server_order_id[..], value);
        self
    }


    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_number_of_days(&mut self, value: i32) -> &mut Self {
        self.number_of_days = value.to_le();
        self
    }


    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.trade_account[..], value);
        self
    }


    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        self.start_date_time = value.to_le();
        self
    }

}

/// This is a message from the Client to the Server to request order fills/executions
/// for an order or orders.
impl HistoricalOrderFillsRequest for HistoricalOrderFillsRequestFixedUnsafe {
    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn server_order_id(&self) -> &str {
        if self.is_out_of_bounds(40) {
            ""
        } else {
            crate::get_fixed(&self.server_order_id[..])
        }
    }

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn number_of_days(&self) -> i32 {
        if self.is_out_of_bounds(44) {
            0
        } else {
            i32::from_le(self.number_of_days)
        }
    }

    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(76) {
            ""
        } else {
            crate::get_fixed(&self.trade_account[..])
        }
    }

    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn start_date_time(&self) -> DateTime {
        if self.is_out_of_bounds(88) {
            0
        } else {
            i64::from_le(self.start_date_time)
        }
    }

    /// A unique request identifier. The Server will return the same identifier
    /// in the response.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }


    /// Leave empty if want all order fills. Otherwise, request order fills for
    /// given Server Order identifier.
    fn set_server_order_id(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(40) {
            crate::set_fixed(&mut self.server_order_id[..], value);
        }
        self
    }


    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The NumberOfDays field specifies to the Server to return order fills counting
    /// from the current day back by the specified number of days.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_number_of_days(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(44) {
            self.number_of_days = value.to_le();
        }
        self
    }


    /// This specifies the particular Trade Account to request order fills for.
    /// This specifies the particular Trade Account to request order fills for.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(76) {
            crate::set_fixed(&mut self.trade_account[..], value);
        }
        self
    }


    /// The NumberOfDays field is ignored by the Server when StartDateTime is
    /// set.
    ///
    /// The StartDateTime field specifies to the Server to return order fills
    /// starting with date time specified.
    ///
    /// If NumberOfDays and StartDateTime are both not set or 0, the Server will
    /// return all historical order fills available.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        if !self.is_out_of_bounds(88) {
            self.start_date_time = value.to_le();
        }
        self
    }

}

