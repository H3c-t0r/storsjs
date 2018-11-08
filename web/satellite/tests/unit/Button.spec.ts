import { shallowMount, mount } from '@vue/test-utils';
import Button from "@/components/Button.vue";
import * as sinon from 'sinon';

describe('Button.vue', () => {
	it('renders correctly with size and label props', () => {
		let label = "testLabel";
		let width = "30px";
		let height = "20px";

		const wrapper = shallowMount(Button, {
			propsData: { label, width, height },
		});

		expect(wrapper.element.style.width).toMatch(width);
		expect(wrapper.element.style.height).toMatch(height);
		expect(wrapper.text()).toMatch(label);
	});

	it('renders correctly with default props', () => {

		const wrapper = shallowMount(Button);

		expect(wrapper.element.style.width).toMatch("inherit");
		expect(wrapper.element.style.height).toMatch("inherit");
		expect(wrapper.text()).toMatch("Default");
	});

	it('trigger onPress correctly', () => {
		let onPressSpy = sinon.spy();

		const wrapper = mount(Button, {
			propsData: {
				onPress: onPressSpy
			}
		});

		wrapper.find('div.container').trigger('click');
		
		expect(onPressSpy.callCount).toBe(1);
	});
});