Videos
Generate videos.

Create video
post
 
https://api.openai.com/v1/videos
Create a video

Request body
prompt
string

Required
Text prompt that describes the video to generate.

input_reference
file

Optional
Optional image reference that guides generation.

model
string

Optional
The video generation model to use. Defaults to sora-2.

seconds
string

Optional
Clip duration in seconds. Defaults to 4 seconds.

size
string

Optional
Output resolution formatted as width x height. Defaults to 720x1280.

Returns
Returns the newly created video job.

Example request
curl https://api.openai.com/v1/videos \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -F "model=sora-2" \
  -F "prompt=A calico cat playing a piano on stage"
Response
{
  "id": "video_123",
  "object": "video",
  "model": "sora-2",
  "status": "queued",
  "progress": 0,
  "created_at": 1712697600,
  "size": "1024x1808",
  "seconds": "8",
  "quality": "standard"
}
Remix video
post
 
https://api.openai.com/v1/videos/{video_id}/remix
Create a video remix

Path parameters
video_id
string

Required
The identifier of the completed video to remix.

Request body
prompt
string

Required
Updated text prompt that directs the remix generation.

Returns
Creates a remix of the specified video job using the provided prompt.

Example request
curl -X POST https://api.openai.com/v1/videos/video_123/remix \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "Extend the scene with the cat taking a bow to the cheering audience"
  }'
Response
{
  "id": "video_456",
  "object": "video",
  "model": "sora-2",
  "status": "queued",
  "progress": 0,
  "created_at": 1712698600,
  "size": "720x1280",
  "seconds": "8",
  "remixed_from_video_id": "video_123"
}
List videos
get
 
https://api.openai.com/v1/videos
List videos

Query parameters
after
string

Optional
Identifier for the last item from the previous pagination request

limit
integer

Optional
Number of items to retrieve

order
string

Optional
Sort order of results by timestamp. Use asc for ascending order or desc for descending order.

Returns
Returns a paginated list of video jobs for the organization.

Example request
curl https://api.openai.com/v1/videos \
  -H "Authorization: Bearer $OPENAI_API_KEY"
Response
{
  "data": [
    {
      "id": "video_123",
      "object": "video",
      "model": "sora-2",
      "status": "completed"
    }
  ],
  "object": "list"
}
Retrieve video
get
 
https://api.openai.com/v1/videos/{video_id}
Retrieve a video

Path parameters
video_id
string

Required
The identifier of the video to retrieve.

Returns
Returns the video job matching the provided identifier.

Example request
import OpenAI from 'openai';

const client = new OpenAI();

const video = await client.videos.retrieve('video_123');

console.log(video.id);
Delete video
delete
 
https://api.openai.com/v1/videos/{video_id}
Delete a video

Path parameters
video_id
string

Required
The identifier of the video to delete.

Returns
Returns the deleted video job metadata.

Example request
import OpenAI from 'openai';

const client = new OpenAI();

const video = await client.videos.delete('video_123');

console.log(video.id);
Retrieve video content
get
 
https://api.openai.com/v1/videos/{video_id}/content
Download video content

Path parameters
video_id
string

Required
The identifier of the video whose media to download.

Query parameters
variant
string

Optional
Which downloadable asset to return. Defaults to the MP4 video.

Returns
Streams the rendered video content for the specified video job.

Example request
import OpenAI from 'openai';

const client = new OpenAI();

const response = await client.videos.downloadContent('video_123');

console.log(response);

const content = await response.blob();
console.log(content);


==================================
Video job
Structured information describing a generated video job.

completed_at
integer

Unix timestamp (seconds) for when the job completed, if finished.

created_at
integer

Unix timestamp (seconds) for when the job was created.

error
object

Error payload that explains why generation failed, if applicable.


Show properties
expires_at
integer

Unix timestamp (seconds) for when the downloadable assets expire, if set.

id
string

Unique identifier for the video job.

model
string

The video generation model that produced the job.

object
string

The object type, which is always video.

progress
integer

Approximate completion percentage for the generation task.

prompt
string

The prompt that was used to generate the video.

remixed_from_video_id
string

Identifier of the source video if this video is a remix.

seconds
string

Duration of the generated clip in seconds.

size
string

The resolution of the generated video.

status
string

Current lifecycle status of the video job.


Video job
Structured information describing a generated video job.

completed_at
integer

Unix timestamp (seconds) for when the job completed, if finished.

created_at
integer

Unix timestamp (seconds) for when the job was created.

error
object

Error payload that explains why generation failed, if applicable.


Hide properties
code
string

message
string

expires_at
integer

Unix timestamp (seconds) for when the downloadable assets expire, if set.

id
string

Unique identifier for the video job.

model
string

The video generation model that produced the job.

object
string

The object type, which is always video.

progress
integer

Approximate completion percentage for the generation task.

prompt
string

The prompt that was used to generate the video.

remixed_from_video_id
string

Identifier of the source video if this video is a remix.

seconds
string

Duration of the generated clip in seconds.

size
string

The resolution of the generated video.

status
string

Current lifecycle status of the video job.