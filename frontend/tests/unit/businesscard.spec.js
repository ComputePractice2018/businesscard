import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import businesscard from '@/components/businesscard.vue'

describe('businesscard.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(businesscard, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders titles', () => {
    const wrapper = shallowMount(businesscard, {})
    expect(wrapper.text()).to.include('ФИО')
    expect(wrapper.text()).to.include('Должность')
    expect(wrapper.text()).to.include('Телефон')
    expect(wrapper.text()).to.include('E-mail')
    expect(wrapper.text()).to.include('Место работы')
    expect(wrapper.text()).to.include('Адрес')
  })
})
