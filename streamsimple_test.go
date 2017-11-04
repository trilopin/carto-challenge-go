package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumFloatField(t *testing.T) {
	str := `VendorID,tpep_pickup_datetime,tpep_dropoff_datetime,passenger_count,trip_distance,pickup_longitude,pickup_latitude,RatecodeID,store_and_fwd_flag,dropoff_longitude,dropoff_latitude,payment_type,fare_amount,extra,mta_tax,tip_amount,tolls_amount,improvement_surcharge,total_amount
2,2016-01-01 00:00:00,2016-01-01 00:00:00,2,1.10,-73.990371704101563,40.734695434570313,1,N,-73.981842041015625,40.732406616210937,2,7.5,0.5,0.5,0,0,0.3,8.8
2,2016-01-01 00:00:00,2016-01-01 00:00:00,5,4.90,-73.980781555175781,40.729911804199219,1,N,-73.944473266601563,40.716678619384766,1,18,0.5,0.5,0,0,0.3,19.3
2,2016-01-01 00:00:00,2016-01-01 00:00:00,1,10.54,-73.984550476074219,40.6795654296875,1,N,-73.950271606445313,40.788925170898438,1,33,0.5,0.5,0,0,0.3,34.3
2,2016-01-01 00:00:00,2016-01-01 00:00:00,1,4.75,-73.99346923828125,40.718990325927734,1,N,-73.962242126464844,40.657333374023437,2,16.5,0,0.5,0,0,0.3,17.3
2,2016-01-01 00:00:00,2016-01-01 00:00:00,3,1.76,-73.960624694824219,40.781330108642578,1,N,-73.977264404296875,40.758514404296875,2,8,0,0.5,0,0,0.3,8.8
2,2016-01-01 00:00:00,2016-01-01 00:18:30,2,5.52,-73.980117797851563,40.743049621582031,1,N,-73.913490295410156,40.763141632080078,2,19,0.5,0.5,0,0,0.3,20.3
2,2016-01-01 00:00:00,2016-01-01 00:26:45,2,7.45,-73.994056701660156,40.719989776611328,1,N,-73.966361999511719,40.789871215820313,2,26,0.5,0.5,0,0,0.3,27.3
1,2016-01-01 00:00:01,2016-01-01 00:11:55,1,1.20,-73.979423522949219,40.744613647460938,1,N,-73.992034912109375,40.753944396972656,2,9,0.5,0.5,0,0,0.3,10.3
1,2016-01-01 00:00:02,2016-01-01 00:11:14,1,6.00,-73.947151184082031,40.791046142578125,1,N,-73.920768737792969,40.865577697753906,2,18,0.5,0.5,0,0,0.3,19.3
2,2016-01-01 00:00:02,2016-01-01 00:11:08,1,3.21,-73.998344421386719,40.723896026611328,1,N,-73.995849609375,40.688400268554688,2,11.5,0.5,0.5,0,0,0.3,12.8
2,2016-01-01 00:00:03,2016-01-01 00:06:19,1,.79,-74.006149291992188,40.744918823242188,1,N,-73.993797302246094,40.741439819335938,2,6,0.5,0.5,0,0,0.3,7.3
2,2016-01-01 00:00:03,2016-01-01 00:15:49,6,2.43,-73.969329833984375,40.763538360595703,1,N,-73.995689392089844,40.744251251220703,1,12,0.5,0.5,3.99,0,0.3,17.29
2,2016-01-01 00:00:03,2016-01-01 00:00:11,4,.01,-73.989021301269531,40.721538543701172,1,N,-73.988960266113281,40.721698760986328,2,2.5,0.5,0.5,0,0,0.3,3.8
1,2016-01-01 00:00:04,2016-01-01 00:14:32,1,3.70,-74.004302978515625,40.742240905761719,1,N,-74.007362365722656,40.706935882568359,1,14,0.5,0.5,3.05,0,0.3,18.35
1,2016-01-01 00:00:05,2016-01-01 00:14:27,2,2.20,-73.991996765136719,40.718578338623047,1,N,-74.005134582519531,40.739944458007813,1,11,0.5,0.5,1.5,0,0.3,13.8
2,2016-01-01 00:00:05,2016-01-01 00:07:17,1,.54,-73.985160827636719,40.768951416015625,1,N,-73.990226745605469,40.761730194091797,2,6,0.5,0.5,0,0,0.3,7.3
2,2016-01-01 00:00:05,2016-01-01 00:07:14,1,1.92,-73.973091125488281,40.795360565185547,1,N,-73.978370666503906,40.773151397705078,2,7.5,0.5,0.5,0,0,0.3,8.8
1,2016-01-01 00:00:06,2016-01-01 00:04:44,1,1.70,-73.982101440429688,40.774696350097656,1,Y,-73.970939636230469,40.796707153320313,1,7,0.5,0.5,1.65,0,0.3,9.95
2,2016-01-01 00:00:06,2016-01-01 00:07:14,1,1.38,-73.994842529296875,40.718498229980469,1,N,-73.98980712890625,40.734230041503906,1,7,0.5,0.5,1.66,0,0.3,9.96
1,2016-01-01 00:00:07,2016-01-01 00:20:35,2,4.90,-73.953033447265625,40.672115325927734,1,N,-73.986572265625,40.710594177246094,1,19,0.5,0.5,4.06,0,0.3,24.36
1,2016-01-01 00:00:07,2016-01-01 00:09:49,1,1.80,-73.989166259765625,40.726589202880859,1,N,-74.009483337402344,40.715072631835938,2,9,0.5,0.5,0,0,0.3,10.3
2,2016-01-01 00:00:08,2016-01-01 00:18:51,1,3.09,-73.999069213867188,40.720172882080078,1,N,-73.973388671875,40.756561279296875,2,14.5,0.5,0.5,0,0,0.3,15.8
2,2016-01-01 00:00:08,2016-01-01 00:04:37,1,.72,-73.997138977050781,40.747219085693359,1,N,-74.004486083984375,40.751796722412109,2,5,0.5,0.5,0,0,0.3,6.3
2,2016-01-01 00:00:08,2016-01-01 00:03:24,1,.69,-73.997413635253906,40.736675262451172,1,N,-73.985664367675781,40.732681274414063,2,4.5,0.5,0.5,0,0,0.3,5.8
1,2016-01-01 00:00:09,2016-01-01 00:19:03,3,5.30,-73.99713134765625,40.736961364746094,1,N,-73.928421020507813,40.755580902099609,1,18,0.5,0.5,3.85,0,0.3,23.15`

	r := strings.NewReader(str)
	count, sum := sumFloatField(r, '\n', 15)
	assert.Equal(t, count, 24)
	assert.Equal(t, sum, 15.91)

}