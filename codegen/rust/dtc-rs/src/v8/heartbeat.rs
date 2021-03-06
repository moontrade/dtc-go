// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const HEARTBEAT_FIXED_SIZE: usize = 16;

/// size                  u16       = HeartbeatFixedSize  (16)
/// type                  u16       = HEARTBEAT  (3)
/// num_dropped_messages  u32       = 0
/// current_date_time     DateTime  = 0
pub(crate) const HEARTBEAT_FIXED_DEFAULT: [u8; 16] =
    [16, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// Both the Client and the Server need to send to the other side a heartbeat
/// at the interval specified by the HeartbeatIntervalInSeconds member in
/// the LogonRequestVLS.
///
/// There are no required member fields to set in this message. The purpose
/// of the HeartbeatFixed message is so that the Client or the Server can
/// determine whether the other side is still connected.
///
/// It is recommended that if there is a loss of HeartbeatFixed messages from
/// the other side, for twice the amount of the HeartbeatIntervalInSeconds
/// time that it is safe to assume that the other side is no longer present
/// and the network socket should be then gracefully closed.
///
/// The Server may choose to send a heartbeat message every second to the
/// Client. In this particular case, it is recommended the Client use a minimum
/// time of about 5 to 10 seconds without a heartbeat to determine the loss
/// of the connection rather than the standard of twice the amount of the
/// heartbeat time interval.
pub trait Heartbeat: Message {
    type Safe: Heartbeat;
    type Unsafe: Heartbeat;

    fn num_dropped_messages(&self) -> u32;

    fn current_date_time(&self) -> DateTime;

    fn set_num_dropped_messages(&mut self, value: u32) -> &mut Self;

    fn set_current_date_time(&mut self, value: DateTime) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl Heartbeat) {
        to.set_num_dropped_messages(self.num_dropped_messages());
        to.set_current_date_time(self.current_date_time());
    }

    #[inline]
    fn parse<F: Fn(Parsed<Self::Safe, Self::Unsafe>) -> Result<(), crate::Error>>(
        data: &[u8],
        f: F,
    ) -> Result<(), crate::Error> {
        if data.len() < 4 {
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

/// Both the Client and the Server need to send to the other side a heartbeat
/// at the interval specified by the HeartbeatIntervalInSeconds member in
/// the LogonRequestVLS.
///
/// There are no required member fields to set in this message. The purpose
/// of the HeartbeatFixed message is so that the Client or the Server can
/// determine whether the other side is still connected.
///
/// It is recommended that if there is a loss of HeartbeatFixed messages from
/// the other side, for twice the amount of the HeartbeatIntervalInSeconds
/// time that it is safe to assume that the other side is no longer present
/// and the network socket should be then gracefully closed.
///
/// The Server may choose to send a heartbeat message every second to the
/// Client. In this particular case, it is recommended the Client use a minimum
/// time of about 5 to 10 seconds without a heartbeat to determine the loss
/// of the connection rather than the standard of twice the amount of the
/// heartbeat time interval.
pub struct HeartbeatFixed {
    data: *const HeartbeatFixedData,
}

pub struct HeartbeatFixedUnsafe {
    data: *const HeartbeatFixedData,
}

#[repr(packed(8), C)]
pub struct HeartbeatFixedData {
    size: u16,
    r#type: u16,
    num_dropped_messages: u32,
    current_date_time: DateTime,
}

impl HeartbeatFixedData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: HEARTBEAT.to_le(),
            num_dropped_messages: 0u32,
            current_date_time: 0i64,
        }
    }
}

unsafe impl Send for HeartbeatFixed {}
unsafe impl Send for HeartbeatFixedUnsafe {}
unsafe impl Send for HeartbeatFixedData {}

impl Drop for HeartbeatFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HeartbeatFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for HeartbeatFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HeartbeatFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for HeartbeatFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HeartbeatFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for HeartbeatFixed {
    type Target = HeartbeatFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HeartbeatFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HeartbeatFixedUnsafe {
    type Target = HeartbeatFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HeartbeatFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for HeartbeatFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HeartbeatFixed(size: {}, type: {}, num_dropped_messages: {}, current_date_time: {})", self.size(), self.r#type(), self.num_dropped_messages(), self.current_date_time()).as_str())
    }
}

impl core::fmt::Debug for HeartbeatFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HeartbeatFixed(size: {}, type: {}, num_dropped_messages: {}, current_date_time: {})", self.size(), self.r#type(), self.num_dropped_messages(), self.current_date_time()).as_str())
    }
}

impl core::fmt::Display for HeartbeatFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HeartbeatFixedUnsafe(size: {}, type: {}, num_dropped_messages: {}, current_date_time: {})", self.size(), self.r#type(), self.num_dropped_messages(), self.current_date_time()).as_str())
    }
}

impl core::fmt::Debug for HeartbeatFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HeartbeatFixedUnsafe(size: {}, type: {}, num_dropped_messages: {}, current_date_time: {})", self.size(), self.r#type(), self.num_dropped_messages(), self.current_date_time()).as_str())
    }
}

impl crate::Message for HeartbeatFixed {
    type Data = HeartbeatFixedData;

    const TYPE: u16 = HEARTBEAT;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HeartbeatFixedData::new()),
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
            data: data as *const HeartbeatFixedData,
        }
    }
}
impl crate::Message for HeartbeatFixedUnsafe {
    type Data = HeartbeatFixedData;

    const TYPE: u16 = HEARTBEAT;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HeartbeatFixedData::new()),
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
            data: data as *const HeartbeatFixedData,
        }
    }
}
/// Both the Client and the Server need to send to the other side a heartbeat
/// at the interval specified by the HeartbeatIntervalInSeconds member in
/// the LogonRequestVLS.
///
/// There are no required member fields to set in this message. The purpose
/// of the HeartbeatFixed message is so that the Client or the Server can
/// determine whether the other side is still connected.
///
/// It is recommended that if there is a loss of HeartbeatFixed messages from
/// the other side, for twice the amount of the HeartbeatIntervalInSeconds
/// time that it is safe to assume that the other side is no longer present
/// and the network socket should be then gracefully closed.
///
/// The Server may choose to send a heartbeat message every second to the
/// Client. In this particular case, it is recommended the Client use a minimum
/// time of about 5 to 10 seconds without a heartbeat to determine the loss
/// of the connection rather than the standard of twice the amount of the
/// heartbeat time interval.
impl Heartbeat for HeartbeatFixed {
    type Safe = HeartbeatFixed;
    type Unsafe = HeartbeatFixedUnsafe;

    fn num_dropped_messages(&self) -> u32 {
        u32::from_le(self.num_dropped_messages)
    }

    fn current_date_time(&self) -> DateTime {
        i64::from_le(self.current_date_time)
    }

    fn set_num_dropped_messages(&mut self, value: u32) -> &mut Self {
        self.num_dropped_messages = value.to_le();
        self
    }

    fn set_current_date_time(&mut self, value: DateTime) -> &mut Self {
        self.current_date_time = value.to_le();
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

/// Both the Client and the Server need to send to the other side a heartbeat
/// at the interval specified by the HeartbeatIntervalInSeconds member in
/// the LogonRequestVLS.
///
/// There are no required member fields to set in this message. The purpose
/// of the HeartbeatFixed message is so that the Client or the Server can
/// determine whether the other side is still connected.
///
/// It is recommended that if there is a loss of HeartbeatFixed messages from
/// the other side, for twice the amount of the HeartbeatIntervalInSeconds
/// time that it is safe to assume that the other side is no longer present
/// and the network socket should be then gracefully closed.
///
/// The Server may choose to send a heartbeat message every second to the
/// Client. In this particular case, it is recommended the Client use a minimum
/// time of about 5 to 10 seconds without a heartbeat to determine the loss
/// of the connection rather than the standard of twice the amount of the
/// heartbeat time interval.
impl Heartbeat for HeartbeatFixedUnsafe {
    type Safe = HeartbeatFixed;
    type Unsafe = HeartbeatFixedUnsafe;

    fn num_dropped_messages(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0u32
        } else {
            u32::from_le(self.num_dropped_messages)
        }
    }

    fn current_date_time(&self) -> DateTime {
        if self.is_out_of_bounds(16) {
            0i64
        } else {
            i64::from_le(self.current_date_time)
        }
    }

    fn set_num_dropped_messages(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.num_dropped_messages = value.to_le();
        }
        self
    }

    fn set_current_date_time(&mut self, value: DateTime) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.current_date_time = value.to_le();
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
                16usize,
                core::mem::size_of::<HeartbeatFixedData>(),
                "HeartbeatFixedData sizeof expected {:} but was {:}",
                16usize,
                core::mem::size_of::<HeartbeatFixedData>()
            );
            assert_eq!(
                16u16,
                HeartbeatFixed::new().size(),
                "HeartbeatFixed sizeof expected {:} but was {:}",
                16u16,
                HeartbeatFixed::new().size(),
            );
            assert_eq!(
                HEARTBEAT,
                HeartbeatFixed::new().r#type(),
                "HeartbeatFixed type expected {:} but was {:}",
                HEARTBEAT,
                HeartbeatFixed::new().r#type(),
            );
            assert_eq!(
                3u16,
                HeartbeatFixed::new().r#type(),
                "HeartbeatFixed type expected {:} but was {:}",
                3u16,
                HeartbeatFixed::new().r#type(),
            );
            let d = HeartbeatFixedData::new();
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
                (core::ptr::addr_of!(d.num_dropped_messages) as usize) - p,
                "num_dropped_messages offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.num_dropped_messages) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.current_date_time) as usize) - p,
                "current_date_time offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.current_date_time) as usize) - p,
            );
        }
    }
}
