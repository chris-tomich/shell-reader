package common

var EscapeTreeRoot *EscapeLeaf = &EscapeLeaf{
	ByteValue: 0x1b,
	Key:       "",
	NextByte:  map[byte]*EscapeLeaf{
		0x4f: {
			ByteValue: 0x4f,
			Key:       "",
			NextByte:  map[byte]*EscapeLeaf{
				0x50: {
					ByteValue: 0x50,
					Key:       "F1",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x51: {
					ByteValue: 0x51,
					Key:       "F2",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x52: {
					ByteValue: 0x52,
					Key:       "F3",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x53: {
					ByteValue: 0x53,
					Key:       "F4",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
			},
		},
		0x5b: {
			ByteValue: 0x5b,
			Key:       "",
			NextByte:  map[byte]*EscapeLeaf{
				0x41: {
					ByteValue: 0x41,
					Key:       "ArrowUp",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x42: {
					ByteValue: 0x42,
					Key:       "ArrowDown",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x31: {
					ByteValue: 0x31,
					Key:       "",
					NextByte:  map[byte]*EscapeLeaf{
						0x35: {
							ByteValue: 0x35,
							Key:       "",
							NextByte:  map[byte]*EscapeLeaf{
								0x7e: {
									ByteValue: 0x7e,
									Key:       "F5",
									NextByte:  map[byte]*EscapeLeaf{
									},
								},
							},
						},
						0x37: {
							ByteValue: 0x37,
							Key:       "",
							NextByte:  map[byte]*EscapeLeaf{
								0x7e: {
									ByteValue: 0x7e,
									Key:       "F6",
									NextByte:  map[byte]*EscapeLeaf{
									},
								},
							},
						},
						0x38: {
							ByteValue: 0x38,
							Key:       "",
							NextByte:  map[byte]*EscapeLeaf{
								0x7e: {
									ByteValue: 0x7e,
									Key:       "F7",
									NextByte:  map[byte]*EscapeLeaf{
									},
								},
							},
						},
						0x39: {
							ByteValue: 0x39,
							Key:       "",
							NextByte:  map[byte]*EscapeLeaf{
								0x7e: {
									ByteValue: 0x7e,
									Key:       "F8",
									NextByte:  map[byte]*EscapeLeaf{
									},
								},
							},
						},
					},
				},
				0x48: {
					ByteValue: 0x48,
					Key:       "Home",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x35: {
					ByteValue: 0x35,
					Key:       "",
					NextByte:  map[byte]*EscapeLeaf{
						0x7e: {
							ByteValue: 0x7e,
							Key:       "PageUp",
							NextByte:  map[byte]*EscapeLeaf{
							},
						},
					},
				},
				0x33: {
					ByteValue: 0x33,
					Key:       "",
					NextByte:  map[byte]*EscapeLeaf{
						0x7e: {
							ByteValue: 0x7e,
							Key:       "Delete",
							NextByte:  map[byte]*EscapeLeaf{
							},
						},
					},
				},
				0x46: {
					ByteValue: 0x46,
					Key:       "End",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x36: {
					ByteValue: 0x36,
					Key:       "",
					NextByte:  map[byte]*EscapeLeaf{
						0x7e: {
							ByteValue: 0x7e,
							Key:       "PageDown",
							NextByte:  map[byte]*EscapeLeaf{
							},
						},
					},
				},
				0x44: {
					ByteValue: 0x44,
					Key:       "ArrowLeft",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x43: {
					ByteValue: 0x43,
					Key:       "ArrowRight",
					NextByte:  map[byte]*EscapeLeaf{
					},
				},
				0x32: {
					ByteValue: 0x32,
					Key:       "",
					NextByte:  map[byte]*EscapeLeaf{
						0x7e: {
							ByteValue: 0x7e,
							Key:       "Insert",
							NextByte:  map[byte]*EscapeLeaf{
							},
						},
						0x30: {
							ByteValue: 0x30,
							Key:       "",
							NextByte:  map[byte]*EscapeLeaf{
								0x7e: {
									ByteValue: 0x7e,
									Key:       "F9",
									NextByte:  map[byte]*EscapeLeaf{
									},
								},
							},
						},
						0x31: {
							ByteValue: 0x31,
							Key:       "",
							NextByte:  map[byte]*EscapeLeaf{
								0x7e: {
									ByteValue: 0x7e,
									Key:       "F10",
									NextByte:  map[byte]*EscapeLeaf{
									},
								},
							},
						},
						0x34: {
							ByteValue: 0x34,
							Key:       "",
							NextByte:  map[byte]*EscapeLeaf{
								0x7e: {
									ByteValue: 0x7e,
									Key:       "F12",
									NextByte:  map[byte]*EscapeLeaf{
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
