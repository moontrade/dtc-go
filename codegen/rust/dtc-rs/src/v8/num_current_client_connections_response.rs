// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE_VLS_SIZE: usize = 22;

pub(crate) const NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE_FIXED_SIZE: usize = 48;

/// size                                   u16     = NumCurrentClientConnectionsResponseVLSSize  (22)
/// type                                   u16     = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
/// base_size                              u16     = NumCurrentClientConnectionsResponseVLSSize  (22)
/// request_id                             u32     = 0
/// username                               string  = ""
/// num_connections_for_different_devices  i32     = 0
/// num_connections_for_same_device        i32     = 0
pub(crate) const NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE_VLS_DEFAULT: [u8; 22] = [
    22, 0, 124, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// size                                   u16       = NumCurrentClientConnectionsResponseFixedSize  (48)
/// type                                   u16       = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
/// request_id                             u32       = 0
/// username                               string32  = ""
/// num_connections_for_different_devices  i32       = 0
/// num_connections_for_same_device        i32       = 0
pub(crate) const NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE_FIXED_DEFAULT: [u8; 48] = [
    48, 0, 124, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait NumCurrentClientConnectionsResponse: Message {
    type Safe: NumCurrentClientConnectionsResponse;
    type Unsafe: NumCurrentClientConnectionsResponse;

    fn request_id(&self) -> u32;

    fn username(&self) -> &str;

    fn num_connections_for_different_devices(&self) -> i32;

    fn num_connections_for_same_device(&self) -> i32;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_username(&mut self, value: &str) -> &mut Self;

    fn set_num_connections_for_different_devices(&mut self, value: i32) -> &mut Self;

    fn set_num_connections_for_same_device(&mut self, value: i32) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl NumCurrentClientConnectionsResponse) {
        to.set_request_id(self.request_id());
        to.set_username(self.username());
        to.set_num_connections_for_different_devices(self.num_connections_for_different_devices());
        to.set_num_connections_for_same_device(self.num_connections_for_same_device());
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

pub struct NumCurrentClientConnectionsResponseVLS {
    data: *const NumCurrentClientConnectionsResponseVLSData,
    capacity: usize,
}

pub struct NumCurrentClientConnectionsResponseVLSUnsafe {
    data: *const NumCurrentClientConnectionsResponseVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct NumCurrentClientConnectionsResponseVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    username: VLS,
    num_connections_for_different_devices: i32,
    num_connections_for_same_device: i32,
}

pub struct NumCurrentClientConnectionsResponseFixed {
    data: *const NumCurrentClientConnectionsResponseFixedData,
}

pub struct NumCurrentClientConnectionsResponseFixedUnsafe {
    data: *const NumCurrentClientConnectionsResponseFixedData,
}

#[repr(packed(1), C)]
pub struct NumCurrentClientConnectionsResponseFixedData {
    size: u16,
    r#type: u16,
    request_id: u32,
    username: [u8; 32],
    num_connections_for_different_devices: i32,
    num_connections_for_same_device: i32,
}

impl NumCurrentClientConnectionsResponseVLSData {
    pub fn new() -> Self {
        Self {
            size: 22u16.to_le(),
            r#type: NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE.to_le(),
            base_size: 22u16.to_le(),
            request_id: 0u32.to_le(),
            username: crate::message::VLS::new(),
            num_connections_for_different_devices: 0i32.to_le(),
            num_connections_for_same_device: 0i32.to_le(),
        }
    }
}

impl NumCurrentClientConnectionsResponseFixedData {
    pub fn new() -> Self {
        Self {
            size: 48u16.to_le(),
            r#type: NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE.to_le(),
            request_id: 0u32.to_le(),
            username: [0; 32],
            num_connections_for_different_devices: 0i32.to_le(),
            num_connections_for_same_device: 0i32.to_le(),
        }
    }
}

unsafe impl Send for NumCurrentClientConnectionsResponseFixed {}
unsafe impl Send for NumCurrentClientConnectionsResponseFixedUnsafe {}
unsafe impl Send for NumCurrentClientConnectionsResponseFixedData {}
unsafe impl Send for NumCurrentClientConnectionsResponseVLS {}
unsafe impl Send for NumCurrentClientConnectionsResponseVLSUnsafe {}
unsafe impl Send for NumCurrentClientConnectionsResponseVLSData {}

impl Drop for NumCurrentClientConnectionsResponseFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for NumCurrentClientConnectionsResponseFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for NumCurrentClientConnectionsResponseVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for NumCurrentClientConnectionsResponseVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for NumCurrentClientConnectionsResponseFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for NumCurrentClientConnectionsResponseFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for NumCurrentClientConnectionsResponseVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for NumCurrentClientConnectionsResponseVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for NumCurrentClientConnectionsResponseFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for NumCurrentClientConnectionsResponseFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for NumCurrentClientConnectionsResponseVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for NumCurrentClientConnectionsResponseVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for NumCurrentClientConnectionsResponseFixed {
    type Target = NumCurrentClientConnectionsResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for NumCurrentClientConnectionsResponseFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for NumCurrentClientConnectionsResponseFixedUnsafe {
    type Target = NumCurrentClientConnectionsResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for NumCurrentClientConnectionsResponseFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for NumCurrentClientConnectionsResponseVLS {
    type Target = NumCurrentClientConnectionsResponseVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for NumCurrentClientConnectionsResponseVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for NumCurrentClientConnectionsResponseVLSUnsafe {
    type Target = NumCurrentClientConnectionsResponseVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for NumCurrentClientConnectionsResponseVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for NumCurrentClientConnectionsResponseFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("NumCurrentClientConnectionsResponseFixed(size: {}, type: {}, request_id: {}, username: \"{}\", num_connections_for_different_devices: {}, num_connections_for_same_device: {})", self.size(), self.r#type(), self.request_id(), self.username(), self.num_connections_for_different_devices(), self.num_connections_for_same_device()).as_str())
    }
}

impl core::fmt::Debug for NumCurrentClientConnectionsResponseFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("NumCurrentClientConnectionsResponseFixed(size: {}, type: {}, request_id: {}, username: \"{}\", num_connections_for_different_devices: {}, num_connections_for_same_device: {})", self.size(), self.r#type(), self.request_id(), self.username(), self.num_connections_for_different_devices(), self.num_connections_for_same_device()).as_str())
    }
}

impl core::fmt::Display for NumCurrentClientConnectionsResponseFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("NumCurrentClientConnectionsResponseFixedUnsafe(size: {}, type: {}, request_id: {}, username: \"{}\", num_connections_for_different_devices: {}, num_connections_for_same_device: {})", self.size(), self.r#type(), self.request_id(), self.username(), self.num_connections_for_different_devices(), self.num_connections_for_same_device()).as_str())
    }
}

impl core::fmt::Debug for NumCurrentClientConnectionsResponseFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("NumCurrentClientConnectionsResponseFixedUnsafe(size: {}, type: {}, request_id: {}, username: \"{}\", num_connections_for_different_devices: {}, num_connections_for_same_device: {})", self.size(), self.r#type(), self.request_id(), self.username(), self.num_connections_for_different_devices(), self.num_connections_for_same_device()).as_str())
    }
}

impl core::fmt::Display for NumCurrentClientConnectionsResponseVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("NumCurrentClientConnectionsResponseVLS(size: {}, type: {}, base_size: {}, request_id: {}, username: \"{}\", num_connections_for_different_devices: {}, num_connections_for_same_device: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.username(), self.num_connections_for_different_devices(), self.num_connections_for_same_device()).as_str())
    }
}

impl core::fmt::Debug for NumCurrentClientConnectionsResponseVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("NumCurrentClientConnectionsResponseVLS(size: {}, type: {}, base_size: {}, request_id: {}, username: \"{}\", num_connections_for_different_devices: {}, num_connections_for_same_device: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.username(), self.num_connections_for_different_devices(), self.num_connections_for_same_device()).as_str())
    }
}

impl core::fmt::Display for NumCurrentClientConnectionsResponseVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("NumCurrentClientConnectionsResponseVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, username: \"{}\", num_connections_for_different_devices: {}, num_connections_for_same_device: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.username(), self.num_connections_for_different_devices(), self.num_connections_for_same_device()).as_str())
    }
}

impl core::fmt::Debug for NumCurrentClientConnectionsResponseVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("NumCurrentClientConnectionsResponseVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, username: \"{}\", num_connections_for_different_devices: {}, num_connections_for_same_device: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.username(), self.num_connections_for_different_devices(), self.num_connections_for_same_device()).as_str())
    }
}

impl crate::Message for NumCurrentClientConnectionsResponseFixed {
    type Data = NumCurrentClientConnectionsResponseFixedData;

    const TYPE: u16 = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE;
    const BASE_SIZE: usize = 48;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                NumCurrentClientConnectionsResponseFixedData::new(),
            ),
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
            data: data as *const NumCurrentClientConnectionsResponseFixedData,
        }
    }
}
impl crate::Message for NumCurrentClientConnectionsResponseFixedUnsafe {
    type Data = NumCurrentClientConnectionsResponseFixedData;

    const TYPE: u16 = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE;
    const BASE_SIZE: usize = 48;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                NumCurrentClientConnectionsResponseFixedData::new(),
            ),
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
            data: data as *const NumCurrentClientConnectionsResponseFixedData,
        }
    }
}
impl crate::Message for NumCurrentClientConnectionsResponseVLS {
    type Data = NumCurrentClientConnectionsResponseVLSData;

    const TYPE: u16 = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE;
    const BASE_SIZE: usize = 22;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                NumCurrentClientConnectionsResponseVLSData::new(),
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
            data: data as *const NumCurrentClientConnectionsResponseVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for NumCurrentClientConnectionsResponseVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const NumCurrentClientConnectionsResponseVLSData;
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
impl crate::Message for NumCurrentClientConnectionsResponseVLSUnsafe {
    type Data = NumCurrentClientConnectionsResponseVLSData;

    const TYPE: u16 = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE;
    const BASE_SIZE: usize = 22;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                NumCurrentClientConnectionsResponseVLSData::new(),
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
            data: data as *const NumCurrentClientConnectionsResponseVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for NumCurrentClientConnectionsResponseVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const NumCurrentClientConnectionsResponseVLSData;
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
impl NumCurrentClientConnectionsResponse for NumCurrentClientConnectionsResponseVLS {
    type Safe = NumCurrentClientConnectionsResponseVLS;
    type Unsafe = NumCurrentClientConnectionsResponseVLSUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn username(&self) -> &str {
        get_vls(self, self.username)
    }

    fn num_connections_for_different_devices(&self) -> i32 {
        i32::from_le(self.num_connections_for_different_devices)
    }

    fn num_connections_for_same_device(&self) -> i32 {
        i32::from_le(self.num_connections_for_same_device)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_username(&mut self, value: &str) -> &mut Self {
        self.username = set_vls(self, self.username, value);
        self
    }

    fn set_num_connections_for_different_devices(&mut self, value: i32) -> &mut Self {
        self.num_connections_for_different_devices = value.to_le();
        self
    }

    fn set_num_connections_for_same_device(&mut self, value: i32) -> &mut Self {
        self.num_connections_for_same_device = value.to_le();
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

impl NumCurrentClientConnectionsResponse for NumCurrentClientConnectionsResponseVLSUnsafe {
    type Safe = NumCurrentClientConnectionsResponseVLS;
    type Unsafe = NumCurrentClientConnectionsResponseVLSUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(10) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn username(&self) -> &str {
        if self.is_out_of_bounds(14) {
            ""
        } else {
            get_vls(self, self.username)
        }
    }

    fn num_connections_for_different_devices(&self) -> i32 {
        if self.is_out_of_bounds(18) {
            0i32.to_le()
        } else {
            i32::from_le(self.num_connections_for_different_devices)
        }
    }

    fn num_connections_for_same_device(&self) -> i32 {
        if self.is_out_of_bounds(22) {
            0i32.to_le()
        } else {
            i32::from_le(self.num_connections_for_same_device)
        }
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_username(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(14) {
            self.username = set_vls(self, self.username, value);
        }
        self
    }

    fn set_num_connections_for_different_devices(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(18) {
            self.num_connections_for_different_devices = value.to_le();
        }
        self
    }

    fn set_num_connections_for_same_device(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(22) {
            self.num_connections_for_same_device = value.to_le();
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

impl NumCurrentClientConnectionsResponse for NumCurrentClientConnectionsResponseFixed {
    type Safe = NumCurrentClientConnectionsResponseFixed;
    type Unsafe = NumCurrentClientConnectionsResponseFixedUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn username(&self) -> &str {
        get_fixed(&self.username[..])
    }

    fn num_connections_for_different_devices(&self) -> i32 {
        i32::from_le(self.num_connections_for_different_devices)
    }

    fn num_connections_for_same_device(&self) -> i32 {
        i32::from_le(self.num_connections_for_same_device)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_username(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.username[..], value);
        self
    }

    fn set_num_connections_for_different_devices(&mut self, value: i32) -> &mut Self {
        self.num_connections_for_different_devices = value.to_le();
        self
    }

    fn set_num_connections_for_same_device(&mut self, value: i32) -> &mut Self {
        self.num_connections_for_same_device = value.to_le();
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

impl NumCurrentClientConnectionsResponse for NumCurrentClientConnectionsResponseFixedUnsafe {
    type Safe = NumCurrentClientConnectionsResponseFixed;
    type Unsafe = NumCurrentClientConnectionsResponseFixedUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn username(&self) -> &str {
        if self.is_out_of_bounds(40) {
            ""
        } else {
            get_fixed(&self.username[..])
        }
    }

    fn num_connections_for_different_devices(&self) -> i32 {
        if self.is_out_of_bounds(44) {
            0i32.to_le()
        } else {
            i32::from_le(self.num_connections_for_different_devices)
        }
    }

    fn num_connections_for_same_device(&self) -> i32 {
        if self.is_out_of_bounds(48) {
            0i32.to_le()
        } else {
            i32::from_le(self.num_connections_for_same_device)
        }
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_username(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(40) {
            set_fixed(&mut self.username[..], value);
        }
        self
    }

    fn set_num_connections_for_different_devices(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(44) {
            self.num_connections_for_different_devices = value.to_le();
        }
        self
    }

    fn set_num_connections_for_same_device(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(48) {
            self.num_connections_for_same_device = value.to_le();
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
                48usize,
                core::mem::size_of::<NumCurrentClientConnectionsResponseFixedData>(),
                "NumCurrentClientConnectionsResponseFixedData sizeof expected {:} but was {:}",
                48usize,
                core::mem::size_of::<NumCurrentClientConnectionsResponseFixedData>()
            );
            assert_eq!(
                48u16,
                NumCurrentClientConnectionsResponseFixed::new().size(),
                "NumCurrentClientConnectionsResponseFixed sizeof expected {:} but was {:}",
                48u16,
                NumCurrentClientConnectionsResponseFixed::new().size(),
            );
            assert_eq!(
                NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE,
                NumCurrentClientConnectionsResponseFixed::new().r#type(),
                "NumCurrentClientConnectionsResponseFixed type expected {:} but was {:}",
                NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE,
                NumCurrentClientConnectionsResponseFixed::new().r#type(),
            );
            assert_eq!(
                10108u16,
                NumCurrentClientConnectionsResponseFixed::new().r#type(),
                "NumCurrentClientConnectionsResponseFixed type expected {:} but was {:}",
                10108u16,
                NumCurrentClientConnectionsResponseFixed::new().r#type(),
            );
            let d = NumCurrentClientConnectionsResponseFixedData::new();
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
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.username) as usize) - p,
                "username offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.username) as usize) - p,
            );
            assert_eq!(
                40usize,
                (core::ptr::addr_of!(d.num_connections_for_different_devices) as usize) - p,
                "num_connections_for_different_devices offset expected {:} but was {:}",
                40usize,
                (core::ptr::addr_of!(d.num_connections_for_different_devices) as usize) - p,
            );
            assert_eq!(
                44usize,
                (core::ptr::addr_of!(d.num_connections_for_same_device) as usize) - p,
                "num_connections_for_same_device offset expected {:} but was {:}",
                44usize,
                (core::ptr::addr_of!(d.num_connections_for_same_device) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                22usize,
                core::mem::size_of::<NumCurrentClientConnectionsResponseVLSData>(),
                "NumCurrentClientConnectionsResponseVLSData sizeof expected {:} but was {:}",
                22usize,
                core::mem::size_of::<NumCurrentClientConnectionsResponseVLSData>()
            );
            assert_eq!(
                22u16,
                NumCurrentClientConnectionsResponseVLS::new().size(),
                "NumCurrentClientConnectionsResponseVLS sizeof expected {:} but was {:}",
                22u16,
                NumCurrentClientConnectionsResponseVLS::new().size(),
            );
            assert_eq!(
                NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE,
                NumCurrentClientConnectionsResponseVLS::new().r#type(),
                "NumCurrentClientConnectionsResponseVLS type expected {:} but was {:}",
                NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE,
                NumCurrentClientConnectionsResponseVLS::new().r#type(),
            );
            assert_eq!(
                10108u16,
                NumCurrentClientConnectionsResponseVLS::new().r#type(),
                "NumCurrentClientConnectionsResponseVLS type expected {:} but was {:}",
                10108u16,
                NumCurrentClientConnectionsResponseVLS::new().r#type(),
            );
            let d = NumCurrentClientConnectionsResponseVLSData::new();
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
                (core::ptr::addr_of!(d.username) as usize) - p,
                "username offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.username) as usize) - p,
            );
            assert_eq!(
                14usize,
                (core::ptr::addr_of!(d.num_connections_for_different_devices) as usize) - p,
                "num_connections_for_different_devices offset expected {:} but was {:}",
                14usize,
                (core::ptr::addr_of!(d.num_connections_for_different_devices) as usize) - p,
            );
            assert_eq!(
                18usize,
                (core::ptr::addr_of!(d.num_connections_for_same_device) as usize) - p,
                "num_connections_for_same_device offset expected {:} but was {:}",
                18usize,
                (core::ptr::addr_of!(d.num_connections_for_same_device) as usize) - p,
            );
        }
    }
}
