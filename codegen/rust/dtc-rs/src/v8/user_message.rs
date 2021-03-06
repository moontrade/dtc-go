// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const USER_MESSAGE_VLS_SIZE: usize = 12;

pub(crate) const USER_MESSAGE_FIXED_SIZE: usize = 262;

/// size              u16     = UserMessageVLSSize  (12)
/// type              u16     = USER_MESSAGE  (700)
/// base_size         u16     = UserMessageVLSSize  (12)
/// user_message      string  = ""
/// is_popup_message  bool    = true
pub(crate) const USER_MESSAGE_VLS_DEFAULT: [u8; 12] = [12, 0, 188, 2, 12, 0, 0, 0, 0, 0, 0, 0];

/// size              u16        = UserMessageFixedSize  (262)
/// type              u16        = USER_MESSAGE  (700)
/// user_message      string256  = ""
/// is_popup_message  bool       = true
pub(crate) const USER_MESSAGE_FIXED_DEFAULT: [u8; 262] = [
    6, 1, 188, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0,
];

/// This message from the Server to the Client is for providing a message
/// to the user.
///
/// This message can be sent even before a LogonResponseVLS.
pub trait UserMessage: Message {
    type Safe: UserMessage;
    type Unsafe: UserMessage;

    /// General message to present to user in the Client.
    fn user_message(&self) -> &str;

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn is_popup_message(&self) -> bool;

    /// General message to present to user in the Client.
    fn set_user_message(&mut self, value: &str) -> &mut Self;

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn set_is_popup_message(&mut self, value: bool) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl UserMessage) {
        to.set_user_message(self.user_message());
        to.set_is_popup_message(self.is_popup_message());
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

/// This message from the Server to the Client is for providing a message
/// to the user.
///
/// This message can be sent even before a LogonResponseVLS.
pub struct UserMessageVLS {
    data: *const UserMessageVLSData,
    capacity: usize,
}

pub struct UserMessageVLSUnsafe {
    data: *const UserMessageVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct UserMessageVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    user_message: VLS,
    is_popup_message: bool,
}

/// This message from the Server to the Client is for providing a message
/// to the user.
///
/// This message can be sent even before a LogonResponseVLS.
pub struct UserMessageFixed {
    data: *const UserMessageFixedData,
}

pub struct UserMessageFixedUnsafe {
    data: *const UserMessageFixedData,
}

#[repr(packed(8), C)]
pub struct UserMessageFixedData {
    size: u16,
    r#type: u16,
    user_message: [u8; 256],
    is_popup_message: bool,
}

impl UserMessageVLSData {
    pub fn new() -> Self {
        Self {
            size: 12u16.to_le(),
            r#type: USER_MESSAGE.to_le(),
            base_size: 12u16.to_le(),
            user_message: crate::message::VLS::new(),
            is_popup_message: true,
        }
    }
}

impl UserMessageFixedData {
    pub fn new() -> Self {
        Self {
            size: 262u16.to_le(),
            r#type: USER_MESSAGE.to_le(),
            user_message: [0; 256],
            is_popup_message: true,
        }
    }
}

unsafe impl Send for UserMessageFixed {}
unsafe impl Send for UserMessageFixedUnsafe {}
unsafe impl Send for UserMessageFixedData {}
unsafe impl Send for UserMessageVLS {}
unsafe impl Send for UserMessageVLSUnsafe {}
unsafe impl Send for UserMessageVLSData {}

impl Drop for UserMessageFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for UserMessageFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for UserMessageVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for UserMessageVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for UserMessageFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for UserMessageFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for UserMessageVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for UserMessageVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for UserMessageFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for UserMessageFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for UserMessageVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for UserMessageVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for UserMessageFixed {
    type Target = UserMessageFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for UserMessageFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for UserMessageFixedUnsafe {
    type Target = UserMessageFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for UserMessageFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for UserMessageVLS {
    type Target = UserMessageVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for UserMessageVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for UserMessageVLSUnsafe {
    type Target = UserMessageVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for UserMessageVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for UserMessageFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(
            format!(
                "UserMessageFixed(size: {}, type: {}, user_message: \"{}\", is_popup_message: {})",
                self.size(),
                self.r#type(),
                self.user_message(),
                self.is_popup_message()
            )
            .as_str(),
        )
    }
}

impl core::fmt::Debug for UserMessageFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(
            format!(
                "UserMessageFixed(size: {}, type: {}, user_message: \"{}\", is_popup_message: {})",
                self.size(),
                self.r#type(),
                self.user_message(),
                self.is_popup_message()
            )
            .as_str(),
        )
    }
}

impl core::fmt::Display for UserMessageFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("UserMessageFixedUnsafe(size: {}, type: {}, user_message: \"{}\", is_popup_message: {})", self.size(), self.r#type(), self.user_message(), self.is_popup_message()).as_str())
    }
}

impl core::fmt::Debug for UserMessageFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("UserMessageFixedUnsafe(size: {}, type: {}, user_message: \"{}\", is_popup_message: {})", self.size(), self.r#type(), self.user_message(), self.is_popup_message()).as_str())
    }
}

impl core::fmt::Display for UserMessageVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("UserMessageVLS(size: {}, type: {}, base_size: {}, user_message: \"{}\", is_popup_message: {})", self.size(), self.r#type(), self.base_size(), self.user_message(), self.is_popup_message()).as_str())
    }
}

impl core::fmt::Debug for UserMessageVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("UserMessageVLS(size: {}, type: {}, base_size: {}, user_message: \"{}\", is_popup_message: {})", self.size(), self.r#type(), self.base_size(), self.user_message(), self.is_popup_message()).as_str())
    }
}

impl core::fmt::Display for UserMessageVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("UserMessageVLSUnsafe(size: {}, type: {}, base_size: {}, user_message: \"{}\", is_popup_message: {})", self.size(), self.r#type(), self.base_size(), self.user_message(), self.is_popup_message()).as_str())
    }
}

impl core::fmt::Debug for UserMessageVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("UserMessageVLSUnsafe(size: {}, type: {}, base_size: {}, user_message: \"{}\", is_popup_message: {})", self.size(), self.r#type(), self.base_size(), self.user_message(), self.is_popup_message()).as_str())
    }
}

impl crate::Message for UserMessageFixed {
    type Data = UserMessageFixedData;

    const TYPE: u16 = USER_MESSAGE;
    const BASE_SIZE: usize = 262;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, UserMessageFixedData::new()),
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
            data: data as *const UserMessageFixedData,
        }
    }
}
impl crate::Message for UserMessageFixedUnsafe {
    type Data = UserMessageFixedData;

    const TYPE: u16 = USER_MESSAGE;
    const BASE_SIZE: usize = 262;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, UserMessageFixedData::new()),
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
            data: data as *const UserMessageFixedData,
        }
    }
}
impl crate::Message for UserMessageVLS {
    type Data = UserMessageVLSData;

    const TYPE: u16 = USER_MESSAGE;
    const BASE_SIZE: usize = 12;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, UserMessageVLSData::new()),
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
            data: data as *const UserMessageVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for UserMessageVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const UserMessageVLSData;
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
impl crate::Message for UserMessageVLSUnsafe {
    type Data = UserMessageVLSData;

    const TYPE: u16 = USER_MESSAGE;
    const BASE_SIZE: usize = 12;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, UserMessageVLSData::new()),
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
            data: data as *const UserMessageVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for UserMessageVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const UserMessageVLSData;
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
/// This message from the Server to the Client is for providing a message
/// to the user.
///
/// This message can be sent even before a LogonResponseVLS.
impl UserMessage for UserMessageVLS {
    type Safe = UserMessageVLS;
    type Unsafe = UserMessageVLSUnsafe;

    /// General message to present to user in the Client.
    fn user_message(&self) -> &str {
        get_vls(self, self.user_message)
    }

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn is_popup_message(&self) -> bool {
        self.is_popup_message
    }

    /// General message to present to user in the Client.
    fn set_user_message(&mut self, value: &str) -> &mut Self {
        self.user_message = set_vls(self, self.user_message, value);
        self
    }

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn set_is_popup_message(&mut self, value: bool) -> &mut Self {
        self.is_popup_message = value;
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

/// This message from the Server to the Client is for providing a message
/// to the user.
///
/// This message can be sent even before a LogonResponseVLS.
impl UserMessage for UserMessageVLSUnsafe {
    type Safe = UserMessageVLS;
    type Unsafe = UserMessageVLSUnsafe;

    /// General message to present to user in the Client.
    fn user_message(&self) -> &str {
        if self.is_out_of_bounds(10) {
            ""
        } else {
            get_vls(self, self.user_message)
        }
    }

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn is_popup_message(&self) -> bool {
        if self.is_out_of_bounds(11) {
            true
        } else {
            self.is_popup_message
        }
    }

    /// General message to present to user in the Client.
    fn set_user_message(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.user_message = set_vls(self, self.user_message, value);
        }
        self
    }

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn set_is_popup_message(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(11) {
            self.is_popup_message = value;
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

/// This message from the Server to the Client is for providing a message
/// to the user.
///
/// This message can be sent even before a LogonResponseVLS.
impl UserMessage for UserMessageFixed {
    type Safe = UserMessageFixed;
    type Unsafe = UserMessageFixedUnsafe;

    /// General message to present to user in the Client.
    fn user_message(&self) -> &str {
        get_fixed(&self.user_message[..])
    }

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn is_popup_message(&self) -> bool {
        self.is_popup_message
    }

    /// General message to present to user in the Client.
    fn set_user_message(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.user_message[..], value);
        self
    }

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn set_is_popup_message(&mut self, value: bool) -> &mut Self {
        self.is_popup_message = value;
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

/// This message from the Server to the Client is for providing a message
/// to the user.
///
/// This message can be sent even before a LogonResponseVLS.
impl UserMessage for UserMessageFixedUnsafe {
    type Safe = UserMessageFixed;
    type Unsafe = UserMessageFixedUnsafe;

    /// General message to present to user in the Client.
    fn user_message(&self) -> &str {
        if self.is_out_of_bounds(260) {
            ""
        } else {
            get_fixed(&self.user_message[..])
        }
    }

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn is_popup_message(&self) -> bool {
        if self.is_out_of_bounds(261) {
            true
        } else {
            self.is_popup_message
        }
    }

    /// General message to present to user in the Client.
    fn set_user_message(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(260) {
            set_fixed(&mut self.user_message[..], value);
        }
        self
    }

    /// The default for this is 1 which signifies that the Server would like the
    /// Client to present the message to the user in a way which will get their
    /// attention. Otherwise, set this to 0 to give the message lower priority
    /// (just add to a log).
    fn set_is_popup_message(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(261) {
            self.is_popup_message = value;
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
                262usize,
                core::mem::size_of::<UserMessageFixedData>(),
                "UserMessageFixedData sizeof expected {:} but was {:}",
                262usize,
                core::mem::size_of::<UserMessageFixedData>()
            );
            assert_eq!(
                262u16,
                UserMessageFixed::new().size(),
                "UserMessageFixed sizeof expected {:} but was {:}",
                262u16,
                UserMessageFixed::new().size(),
            );
            assert_eq!(
                USER_MESSAGE,
                UserMessageFixed::new().r#type(),
                "UserMessageFixed type expected {:} but was {:}",
                USER_MESSAGE,
                UserMessageFixed::new().r#type(),
            );
            assert_eq!(
                700u16,
                UserMessageFixed::new().r#type(),
                "UserMessageFixed type expected {:} but was {:}",
                700u16,
                UserMessageFixed::new().r#type(),
            );
            let d = UserMessageFixedData::new();
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
                (core::ptr::addr_of!(d.user_message) as usize) - p,
                "user_message offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.user_message) as usize) - p,
            );
            assert_eq!(
                260usize,
                (core::ptr::addr_of!(d.is_popup_message) as usize) - p,
                "is_popup_message offset expected {:} but was {:}",
                260usize,
                (core::ptr::addr_of!(d.is_popup_message) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                12usize,
                core::mem::size_of::<UserMessageVLSData>(),
                "UserMessageVLSData sizeof expected {:} but was {:}",
                12usize,
                core::mem::size_of::<UserMessageVLSData>()
            );
            assert_eq!(
                12u16,
                UserMessageVLS::new().size(),
                "UserMessageVLS sizeof expected {:} but was {:}",
                12u16,
                UserMessageVLS::new().size(),
            );
            assert_eq!(
                USER_MESSAGE,
                UserMessageVLS::new().r#type(),
                "UserMessageVLS type expected {:} but was {:}",
                USER_MESSAGE,
                UserMessageVLS::new().r#type(),
            );
            assert_eq!(
                700u16,
                UserMessageVLS::new().r#type(),
                "UserMessageVLS type expected {:} but was {:}",
                700u16,
                UserMessageVLS::new().r#type(),
            );
            let d = UserMessageVLSData::new();
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
                (core::ptr::addr_of!(d.user_message) as usize) - p,
                "user_message offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.user_message) as usize) - p,
            );
            assert_eq!(
                10usize,
                (core::ptr::addr_of!(d.is_popup_message) as usize) - p,
                "is_popup_message offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.is_popup_message) as usize) - p,
            );
        }
    }
}
