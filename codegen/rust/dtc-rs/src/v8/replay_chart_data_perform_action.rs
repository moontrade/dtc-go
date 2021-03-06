// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const REPLAY_CHART_DATA_PERFORM_ACTION_VLS_SIZE: usize = 18;

pub(crate) const REPLAY_CHART_DATA_PERFORM_ACTION_FIXED_SIZE: usize = 16;

/// size          u16                        = ReplayChartDataPerformActionVLSSize  (18)
/// type          u16                        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
/// base_size     u16                        = ReplayChartDataPerformActionVLSSize  (18)
/// request_id    u32                        = 0
/// action        ReplayChartDataActionEnum  = REPLAY_CHART_DATA_ACTION_NONE  (0)
/// replay_speed  f32                        = 0
pub(crate) const REPLAY_CHART_DATA_PERFORM_ACTION_VLS_DEFAULT: [u8; 18] =
    [18, 0, 121, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size          u16                        = ReplayChartDataPerformActionFixedSize  (16)
/// type          u16                        = REPLAY_CHART_DATA_PERFORM_ACTION  (10105)
/// request_id    u32                        = 0
/// action        ReplayChartDataActionEnum  = REPLAY_CHART_DATA_ACTION_NONE  (0)
/// replay_speed  f32                        = 0
pub(crate) const REPLAY_CHART_DATA_PERFORM_ACTION_FIXED_DEFAULT: [u8; 16] =
    [16, 0, 121, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait ReplayChartDataPerformAction: Message {
    type Safe: ReplayChartDataPerformAction;
    type Unsafe: ReplayChartDataPerformAction;

    fn request_id(&self) -> u32;

    fn action(&self) -> ReplayChartDataActionEnum;

    fn replay_speed(&self) -> f32;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_action(&mut self, value: ReplayChartDataActionEnum) -> &mut Self;

    fn set_replay_speed(&mut self, value: f32) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl ReplayChartDataPerformAction) {
        to.set_request_id(self.request_id());
        to.set_action(self.action());
        to.set_replay_speed(self.replay_speed());
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

pub struct ReplayChartDataPerformActionVLS {
    data: *const ReplayChartDataPerformActionVLSData,
    capacity: usize,
}

pub struct ReplayChartDataPerformActionVLSUnsafe {
    data: *const ReplayChartDataPerformActionVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct ReplayChartDataPerformActionVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    action: ReplayChartDataActionEnum,
    replay_speed: f32,
}

pub struct ReplayChartDataPerformActionFixed {
    data: *const ReplayChartDataPerformActionFixedData,
}

pub struct ReplayChartDataPerformActionFixedUnsafe {
    data: *const ReplayChartDataPerformActionFixedData,
}

#[repr(packed(1), C)]
pub struct ReplayChartDataPerformActionFixedData {
    size: u16,
    r#type: u16,
    request_id: u32,
    action: ReplayChartDataActionEnum,
    replay_speed: f32,
}

impl ReplayChartDataPerformActionVLSData {
    pub fn new() -> Self {
        Self {
            size: 18u16.to_le(),
            r#type: REPLAY_CHART_DATA_PERFORM_ACTION.to_le(),
            base_size: 18u16.to_le(),
            request_id: 0u32.to_le(),
            action: ReplayChartDataActionEnum::ReplayChartDataActionNone.to_le(),
            replay_speed: 0.0,
        }
    }
}

impl ReplayChartDataPerformActionFixedData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: REPLAY_CHART_DATA_PERFORM_ACTION.to_le(),
            request_id: 0u32.to_le(),
            action: ReplayChartDataActionEnum::ReplayChartDataActionNone.to_le(),
            replay_speed: 0.0,
        }
    }
}

unsafe impl Send for ReplayChartDataPerformActionFixed {}
unsafe impl Send for ReplayChartDataPerformActionFixedUnsafe {}
unsafe impl Send for ReplayChartDataPerformActionFixedData {}
unsafe impl Send for ReplayChartDataPerformActionVLS {}
unsafe impl Send for ReplayChartDataPerformActionVLSUnsafe {}
unsafe impl Send for ReplayChartDataPerformActionVLSData {}

impl Drop for ReplayChartDataPerformActionFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for ReplayChartDataPerformActionFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for ReplayChartDataPerformActionVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for ReplayChartDataPerformActionVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for ReplayChartDataPerformActionFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for ReplayChartDataPerformActionFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for ReplayChartDataPerformActionVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for ReplayChartDataPerformActionVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for ReplayChartDataPerformActionFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for ReplayChartDataPerformActionFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for ReplayChartDataPerformActionVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for ReplayChartDataPerformActionVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for ReplayChartDataPerformActionFixed {
    type Target = ReplayChartDataPerformActionFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ReplayChartDataPerformActionFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for ReplayChartDataPerformActionFixedUnsafe {
    type Target = ReplayChartDataPerformActionFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ReplayChartDataPerformActionFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for ReplayChartDataPerformActionVLS {
    type Target = ReplayChartDataPerformActionVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ReplayChartDataPerformActionVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for ReplayChartDataPerformActionVLSUnsafe {
    type Target = ReplayChartDataPerformActionVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ReplayChartDataPerformActionVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for ReplayChartDataPerformActionFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ReplayChartDataPerformActionFixed(size: {}, type: {}, request_id: {}, action: {}, replay_speed: {})", self.size(), self.r#type(), self.request_id(), self.action(), self.replay_speed()).as_str())
    }
}

impl core::fmt::Debug for ReplayChartDataPerformActionFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ReplayChartDataPerformActionFixed(size: {}, type: {}, request_id: {}, action: {}, replay_speed: {})", self.size(), self.r#type(), self.request_id(), self.action(), self.replay_speed()).as_str())
    }
}

impl core::fmt::Display for ReplayChartDataPerformActionFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ReplayChartDataPerformActionFixedUnsafe(size: {}, type: {}, request_id: {}, action: {}, replay_speed: {})", self.size(), self.r#type(), self.request_id(), self.action(), self.replay_speed()).as_str())
    }
}

impl core::fmt::Debug for ReplayChartDataPerformActionFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ReplayChartDataPerformActionFixedUnsafe(size: {}, type: {}, request_id: {}, action: {}, replay_speed: {})", self.size(), self.r#type(), self.request_id(), self.action(), self.replay_speed()).as_str())
    }
}

impl core::fmt::Display for ReplayChartDataPerformActionVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ReplayChartDataPerformActionVLS(size: {}, type: {}, base_size: {}, request_id: {}, action: {}, replay_speed: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.action(), self.replay_speed()).as_str())
    }
}

impl core::fmt::Debug for ReplayChartDataPerformActionVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ReplayChartDataPerformActionVLS(size: {}, type: {}, base_size: {}, request_id: {}, action: {}, replay_speed: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.action(), self.replay_speed()).as_str())
    }
}

impl core::fmt::Display for ReplayChartDataPerformActionVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ReplayChartDataPerformActionVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, action: {}, replay_speed: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.action(), self.replay_speed()).as_str())
    }
}

impl core::fmt::Debug for ReplayChartDataPerformActionVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ReplayChartDataPerformActionVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, action: {}, replay_speed: {})", self.size(), self.r#type(), self.base_size(), self.request_id(), self.action(), self.replay_speed()).as_str())
    }
}

impl crate::Message for ReplayChartDataPerformActionFixed {
    type Data = ReplayChartDataPerformActionFixedData;

    const TYPE: u16 = REPLAY_CHART_DATA_PERFORM_ACTION;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                ReplayChartDataPerformActionFixedData::new(),
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
            data: data as *const ReplayChartDataPerformActionFixedData,
        }
    }
}
impl crate::Message for ReplayChartDataPerformActionFixedUnsafe {
    type Data = ReplayChartDataPerformActionFixedData;

    const TYPE: u16 = REPLAY_CHART_DATA_PERFORM_ACTION;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                ReplayChartDataPerformActionFixedData::new(),
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
            data: data as *const ReplayChartDataPerformActionFixedData,
        }
    }
}
impl crate::Message for ReplayChartDataPerformActionVLS {
    type Data = ReplayChartDataPerformActionVLSData;

    const TYPE: u16 = REPLAY_CHART_DATA_PERFORM_ACTION;
    const BASE_SIZE: usize = 18;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, ReplayChartDataPerformActionVLSData::new()),
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
            data: data as *const ReplayChartDataPerformActionVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for ReplayChartDataPerformActionVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const ReplayChartDataPerformActionVLSData;
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
impl crate::Message for ReplayChartDataPerformActionVLSUnsafe {
    type Data = ReplayChartDataPerformActionVLSData;

    const TYPE: u16 = REPLAY_CHART_DATA_PERFORM_ACTION;
    const BASE_SIZE: usize = 18;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, ReplayChartDataPerformActionVLSData::new()),
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
            data: data as *const ReplayChartDataPerformActionVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for ReplayChartDataPerformActionVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const ReplayChartDataPerformActionVLSData;
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
impl ReplayChartDataPerformAction for ReplayChartDataPerformActionVLS {
    type Safe = ReplayChartDataPerformActionVLS;
    type Unsafe = ReplayChartDataPerformActionVLSUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn action(&self) -> ReplayChartDataActionEnum {
        ReplayChartDataActionEnum::from_le(self.action)
    }

    fn replay_speed(&self) -> f32 {
        f32_le(self.replay_speed)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_action(&mut self, value: ReplayChartDataActionEnum) -> &mut Self {
        self.action = unsafe { core::mem::transmute((value as i32).to_le()) };
        self
    }

    fn set_replay_speed(&mut self, value: f32) -> &mut Self {
        self.replay_speed = f32_le(value);
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

impl ReplayChartDataPerformAction for ReplayChartDataPerformActionVLSUnsafe {
    type Safe = ReplayChartDataPerformActionVLS;
    type Unsafe = ReplayChartDataPerformActionVLSUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(10) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn action(&self) -> ReplayChartDataActionEnum {
        if self.is_out_of_bounds(14) {
            ReplayChartDataActionEnum::ReplayChartDataActionNone.to_le()
        } else {
            ReplayChartDataActionEnum::from_le(self.action)
        }
    }

    fn replay_speed(&self) -> f32 {
        if self.is_out_of_bounds(18) {
            0.0
        } else {
            f32_le(self.replay_speed)
        }
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_action(&mut self, value: ReplayChartDataActionEnum) -> &mut Self {
        if !self.is_out_of_bounds(14) {
            self.action = unsafe { core::mem::transmute((value as i32).to_le()) };
        }
        self
    }

    fn set_replay_speed(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(18) {
            self.replay_speed = f32_le(value);
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

impl ReplayChartDataPerformAction for ReplayChartDataPerformActionFixed {
    type Safe = ReplayChartDataPerformActionFixed;
    type Unsafe = ReplayChartDataPerformActionFixedUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn action(&self) -> ReplayChartDataActionEnum {
        ReplayChartDataActionEnum::from_le(self.action)
    }

    fn replay_speed(&self) -> f32 {
        f32_le(self.replay_speed)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_action(&mut self, value: ReplayChartDataActionEnum) -> &mut Self {
        self.action = unsafe { core::mem::transmute((value as i32).to_le()) };
        self
    }

    fn set_replay_speed(&mut self, value: f32) -> &mut Self {
        self.replay_speed = f32_le(value);
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

impl ReplayChartDataPerformAction for ReplayChartDataPerformActionFixedUnsafe {
    type Safe = ReplayChartDataPerformActionFixed;
    type Unsafe = ReplayChartDataPerformActionFixedUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn action(&self) -> ReplayChartDataActionEnum {
        if self.is_out_of_bounds(12) {
            ReplayChartDataActionEnum::ReplayChartDataActionNone.to_le()
        } else {
            ReplayChartDataActionEnum::from_le(self.action)
        }
    }

    fn replay_speed(&self) -> f32 {
        if self.is_out_of_bounds(16) {
            0.0
        } else {
            f32_le(self.replay_speed)
        }
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_action(&mut self, value: ReplayChartDataActionEnum) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.action = unsafe { core::mem::transmute((value as i32).to_le()) };
        }
        self
    }

    fn set_replay_speed(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.replay_speed = f32_le(value);
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
                core::mem::size_of::<ReplayChartDataPerformActionFixedData>(),
                "ReplayChartDataPerformActionFixedData sizeof expected {:} but was {:}",
                16usize,
                core::mem::size_of::<ReplayChartDataPerformActionFixedData>()
            );
            assert_eq!(
                16u16,
                ReplayChartDataPerformActionFixed::new().size(),
                "ReplayChartDataPerformActionFixed sizeof expected {:} but was {:}",
                16u16,
                ReplayChartDataPerformActionFixed::new().size(),
            );
            assert_eq!(
                REPLAY_CHART_DATA_PERFORM_ACTION,
                ReplayChartDataPerformActionFixed::new().r#type(),
                "ReplayChartDataPerformActionFixed type expected {:} but was {:}",
                REPLAY_CHART_DATA_PERFORM_ACTION,
                ReplayChartDataPerformActionFixed::new().r#type(),
            );
            assert_eq!(
                10105u16,
                ReplayChartDataPerformActionFixed::new().r#type(),
                "ReplayChartDataPerformActionFixed type expected {:} but was {:}",
                10105u16,
                ReplayChartDataPerformActionFixed::new().r#type(),
            );
            let d = ReplayChartDataPerformActionFixedData::new();
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
                (core::ptr::addr_of!(d.action) as usize) - p,
                "action offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.action) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.replay_speed) as usize) - p,
                "replay_speed offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.replay_speed) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                18usize,
                core::mem::size_of::<ReplayChartDataPerformActionVLSData>(),
                "ReplayChartDataPerformActionVLSData sizeof expected {:} but was {:}",
                18usize,
                core::mem::size_of::<ReplayChartDataPerformActionVLSData>()
            );
            assert_eq!(
                18u16,
                ReplayChartDataPerformActionVLS::new().size(),
                "ReplayChartDataPerformActionVLS sizeof expected {:} but was {:}",
                18u16,
                ReplayChartDataPerformActionVLS::new().size(),
            );
            assert_eq!(
                REPLAY_CHART_DATA_PERFORM_ACTION,
                ReplayChartDataPerformActionVLS::new().r#type(),
                "ReplayChartDataPerformActionVLS type expected {:} but was {:}",
                REPLAY_CHART_DATA_PERFORM_ACTION,
                ReplayChartDataPerformActionVLS::new().r#type(),
            );
            assert_eq!(
                10105u16,
                ReplayChartDataPerformActionVLS::new().r#type(),
                "ReplayChartDataPerformActionVLS type expected {:} but was {:}",
                10105u16,
                ReplayChartDataPerformActionVLS::new().r#type(),
            );
            let d = ReplayChartDataPerformActionVLSData::new();
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
                (core::ptr::addr_of!(d.action) as usize) - p,
                "action offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.action) as usize) - p,
            );
            assert_eq!(
                14usize,
                (core::ptr::addr_of!(d.replay_speed) as usize) - p,
                "replay_speed offset expected {:} but was {:}",
                14usize,
                (core::ptr::addr_of!(d.replay_speed) as usize) - p,
            );
        }
    }
}
