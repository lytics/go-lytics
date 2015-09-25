package main

func (c *Cli) getSegments(segments []string) (interface{}, error) {
	if len(segments) == 1 {
		segment, err := c.Client.GetSegment(segments[0])
		if err != nil {
			return nil, err
		}

		return segment, nil
	} else {
		segments, err := c.Client.GetSegments()
		if err != nil {
			return nil, err
		}

		return segments, nil
	}
}

func (c *Cli) getSegmentSizes(segments []string) (interface{}, error) {
	if len(segments) == 1 {
		segment, err := c.Client.GetSegmentSize(segments[0])
		if err != nil {
			return nil, err
		}

		return segment, nil
	} else {
		segments, err := c.Client.GetSegmentSizes(segments)
		if err != nil {
			return nil, err
		}

		return segments, nil
	}
}

func (c *Cli) getSegmentAttributions(segments []string, limit int) (interface{}, error) {
	attributions, err := c.Client.GetSegmentAttribution(segments, limit)
	if err != nil {
		return nil, err
	}

	return attributions, nil
}
