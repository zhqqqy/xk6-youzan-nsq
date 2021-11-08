import {check} from 'k6';
import {Counter, Rate, Trend} from 'k6/metrics';
import {close, consume, received,} from 'k6/x/nsq'; // import mqtt plugin

const lookupAddress = ["10.215.49.141:4161"];
const nsqTopic = "sinker_test";
const nsqChannel = "sinker";

const consumer = consume(lookupAddress, nsqTopic, nsqChannel);

// 新建一个类型为 Counter 名为 my_counter 的自定制指标
let myCounter = new Counter('my_counter');
let myRate = new Rate("my_rate");
//导出一个选项，设置vus(虚拟用户数为 2)
export let options = {
    vus: 2,
    duration: '100s',

};

export default function () {
    let startTime = new Date().getTime();
    let messages = received(lookupAddress, consumer, 10, 3);
    let paased = check(messages, {
        "is content correct": msg => msg != null
    });
    myCounter.add(1)
    myRate.add(paased)
}

export function teardown(data) {
    close(consumer);
}



